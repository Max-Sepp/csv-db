package btree

import (
	"fmt"
	"sort"
	"testing"
)

// a small number of keys so that the fuzzer creates lots of duplicate keys
const fuzzNumKeys = 8

func fuzzKey(b byte) string {
	return fmt.Sprintf("k%d", int(b)%fuzzNumKeys)
}

// FuzzBtree runs a random sequence of operations on Btrees of several orders and compares them to a
// simple model (a map from key to the row pointers stored under that key). After every operation the
// structure of the tree is checked, as well as its contents and the result of Find for every key.
//
// ops is read two bytes at a time: an operation and its argument.
func FuzzBtree(f *testing.F) {
	f.Add([]byte{0, 1, 0, 2, 0, 3, 2, 1, 3, 2})
	f.Add([]byte{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 2, 0, 2, 0, 2, 0, 2, 0, 1, 0, 1})
	f.Add([]byte{0, 0, 0, 1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6, 2, 3, 2, 9, 2, 0})
	// found by the fuzzer: deleting a duplicate key from an internal node removed the wrong entry
	f.Add([]byte("00000202000002200"))

	f.Fuzz(func(t *testing.T, ops []byte) {
		for order := 3; order <= 7; order++ {
			tree := New(order)
			model := map[string][]int64{}
			var nextRowPtr int64

			for i := 0; i+1 < len(ops) && i < 2000; i += 2 {
				op, arg := ops[i], ops[i+1]
				key := fuzzKey(arg)

				switch op % 4 {
				case 0, 1:
					tree.Insert(key, nextRowPtr)
					model[key] = append(model[key], nextRowPtr)
					nextRowPtr++

				case 2:
					rowPtr, err := tree.Delete(key)

					if len(model[key]) == 0 {
						if err != ErrKeyNotFound {
							t.Fatalf("order %d: Delete(%q) of a missing key: got error %v, want ErrKeyNotFound", order, key, err)
						}
						break
					}

					if err != nil {
						t.Fatalf("order %d: Delete(%q): unexpected error %v", order, key, err)
					}

					index := indexOf(model[key], rowPtr)
					if index == -1 {
						t.Fatalf("order %d: Delete(%q) returned row pointer %d which is not stored under that key, stored: %v", order, key, rowPtr, model[key])
					}
					model[key] = append(model[key][:index], model[key][index+1:]...)

				case 3:
					// either delete an entry that exists or one with a row pointer that was never used
					rowPtr := nextRowPtr + 1
					if len(model[key]) > 0 && arg >= 128 {
						rowPtr = model[key][int(arg)%len(model[key])]
					}

					err := tree.DeleteEntry(key, rowPtr)

					index := indexOf(model[key], rowPtr)
					if index == -1 {
						if err != ErrKeyNotFound {
							t.Fatalf("order %d: DeleteEntry(%q, %d) of a missing entry: got error %v, want ErrKeyNotFound", order, key, rowPtr, err)
						}
						break
					}

					if err != nil {
						t.Fatalf("order %d: DeleteEntry(%q, %d): unexpected error %v", order, key, rowPtr, err)
					}
					model[key] = append(model[key][:index], model[key][index+1:]...)
				}

				checkTree(t, tree, model)
			}
		}
	})
}

func indexOf(slice []int64, value int64) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func checkTree(t *testing.T, tree *Btree, model map[string][]int64) {
	t.Helper()

	checkNode(t, tree, tree.root, true, nil, nil)

	// contents must be in order and hold exactly the entries of the model
	array := tree.ToArray()
	got := []string{}
	for i, entry := range array {
		if i > 0 && array[i-1].key > entry.key {
			t.Fatalf("ToArray is not sorted: %q comes before %q", array[i-1].key, entry.key)
		}
		got = append(got, fmt.Sprintf("%s/%d", entry.key, entry.rowPtr))
	}

	want := []string{}
	for key, rowPtrs := range model {
		for _, rowPtr := range rowPtrs {
			want = append(want, fmt.Sprintf("%s/%d", key, rowPtr))
		}
	}

	sort.Strings(got)
	sort.Strings(want)
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("tree contents differ from model\ngot:  %v\nwant: %v", got, want)
	}

	for i := 0; i < fuzzNumKeys; i++ {
		key := fuzzKey(byte(i))
		found, foundKey, rowPtr := tree.Find(key)

		if found != (len(model[key]) > 0) {
			t.Fatalf("Find(%q) found = %v but model has %d entries", key, found, len(model[key]))
		}
		if found && (foundKey != key || indexOf(model[key], rowPtr) == -1) {
			t.Fatalf("Find(%q) returned (%q, %d) which is not in the model: %v", key, foundKey, rowPtr, model[key])
		}
	}
}

// checkNode checks the B-tree invariants of the subtree at treeNode where every key has to be
// between lower and upper (nil means no bound). Returns the height of the subtree.
func checkNode(t *testing.T, tree *Btree, treeNode *node, isRoot bool, lower *string, upper *string) int {
	t.Helper()

	if len(treeNode.keys) > tree.maxElements {
		t.Fatalf("node has %d keys, maximum is %d", len(treeNode.keys), tree.maxElements)
	}
	if !isRoot && len(treeNode.keys) < tree.minElements {
		t.Fatalf("node has %d keys, minimum is %d", len(treeNode.keys), tree.minElements)
	}

	for i, entry := range treeNode.keys {
		if lower != nil && entry.key < *lower || upper != nil && entry.key > *upper {
			t.Fatalf("key %q is outside the range allowed by its parent", entry.key)
		}
		if i > 0 && treeNode.keys[i-1].key > entry.key {
			t.Fatalf("keys in a node are not sorted: %q before %q", treeNode.keys[i-1].key, entry.key)
		}
	}

	if treeNode.leaf {
		if len(treeNode.child) != 0 {
			t.Fatalf("leaf node has %d children", len(treeNode.child))
		}
		return 1
	}

	if len(treeNode.child) != len(treeNode.keys)+1 {
		t.Fatalf("node has %d keys and %d children", len(treeNode.keys), len(treeNode.child))
	}

	height := -1
	for i, child := range treeNode.child {
		childLower, childUpper := lower, upper
		if i > 0 {
			childLower = &treeNode.keys[i-1].key
		}
		if i < len(treeNode.keys) {
			childUpper = &treeNode.keys[i].key
		}

		childHeight := checkNode(t, tree, child, false, childLower, childUpper)
		if height != -1 && childHeight != height {
			t.Fatal("leaves are not all at the same depth")
		}
		height = childHeight
	}
	return height + 1
}
