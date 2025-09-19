package btree

type keyStruct struct {
	key    string
	rowPtr int64
}

type node struct {
	leaf  bool
	keys  []keyStruct
	child []*node
}

type Btree struct {
	root        *node
	maxElements int
	minElements int
}

func New(order int) *Btree {
	tree := new(Btree)
	tree.maxElements = order - 1
	tree.minElements = tree.maxElements / 2

	tree.root = &node{
		leaf:  true,
		keys:  []keyStruct{},
		child: []*node{},
	}
	return tree
}

func (tree *Btree) ToArray() []keyStruct {
	return toArray(tree.root)
}

func toArray(treeNode *node) []keyStruct {
	if treeNode.leaf {
		return treeNode.keys
	} else {
		answer := []keyStruct{}
		answer = append(answer, toArray(treeNode.child[0])...)
		for i := 0; i < len(treeNode.keys); i++ {
			answer = append(answer, treeNode.keys[i])
			answer = append(answer, toArray(treeNode.child[i+1])...)
		}
		return answer
	}
}

func (treeNode *node) findKeyIndex(key string) int {
	i := 0
	for i < len(treeNode.keys) && key > treeNode.keys[i].key {
		i++
	}
	return i
}

func insertIntoSlice[T any](slice []T, place int, item T) []T {
	if len(slice) <= 0 || len(slice) <= place {
		return append(slice, item)
	}
	slice = append(slice[:place+1], slice[place:]...)
	slice[place] = item
	return slice
}

func popFromSlice[T any](slice []T, place int) []T {
	return append(slice[:place], slice[place+1:]...)
}
