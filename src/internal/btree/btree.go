package Btree

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

func New(maxElements int) *Btree {
	tree := new(Btree)
	tree.maxElements = maxElements
	tree.minElements = maxElements / 2

	tree.root = &node{
		leaf:  true,
		keys:  []keyStruct{},
		child: []*node{},
	}
	return tree
}

func (tree *Btree) Insert(keyStr string, rowPtr int64) {
	key := keyStruct{
		key:    keyStr,
		rowPtr: rowPtr,
	}
	rebalanceRequired, middleKey, rightNode := tree.insertHelper(tree.root, key)

	if !rebalanceRequired {
		return
	}

	leftNodeTemp := tree.root
	tree.root = &node{
		leaf:  false,
		keys:  []keyStruct{middleKey},
		child: []*node{leftNodeTemp, rightNode},
	}
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

func (tree *Btree) insertHelper(treeNode *node, key keyStruct) (rebalanceRequired bool, middleKey keyStruct, rightNode *node) {
	if treeNode.leaf {
		// add key to keys of treeNode since it is a leaf treeNode and has space
		i := 0
		for i < len(treeNode.keys) && key.key > treeNode.keys[i].key {
			i++
		}
		treeNode.keys = insertIntoSlice(treeNode.keys, i, key)

		if len(treeNode.keys) <= tree.maxElements {
			return false, keyStruct{}, nil
		}

		// split treeNode
		middleKeyIndex := tree.minElements
		middleKey := treeNode.keys[middleKeyIndex]
		treeNode.keys = popFromSlice(treeNode.keys, middleKeyIndex)

		rightNode := &node{
			leaf:  treeNode.leaf,
			keys:  []keyStruct{},
			child: []*node{},
		}

		for middleKeyIndex < len(treeNode.keys) {
			rightNode.keys = append(rightNode.keys, treeNode.keys[middleKeyIndex])
			treeNode.keys = popFromSlice(treeNode.keys, middleKeyIndex)
		}

		for middleKeyIndex+1 < len(treeNode.child) {
			rightNode.child = append(rightNode.child, treeNode.child[middleKeyIndex])
			treeNode.child = popFromSlice(treeNode.child, middleKeyIndex)
		}

		return true, middleKey, rightNode
	} else {
		i := 0
		for i < len(treeNode.keys) && key.key > treeNode.keys[i].key {
			i++
		}

		rebalanceRequired, middleKey, rightNode := tree.insertHelper(treeNode.child[i], key)

		if !rebalanceRequired {
			return false, keyStruct{}, nil
		}

		// add returned key and right node
		i = 0
		for i < len(treeNode.keys) && middleKey.key > treeNode.keys[i].key {
			i++
		}
		treeNode.keys = insertIntoSlice(treeNode.keys, i, middleKey)
		treeNode.child = insertIntoSlice(treeNode.child, i+1, rightNode)

		if len(treeNode.keys) <= tree.maxElements {
			return false, keyStruct{}, nil
		}

		// split treeNode
		middleKeyIndex := tree.minElements
		middleKey = treeNode.keys[middleKeyIndex]
		treeNode.keys = popFromSlice(treeNode.keys, middleKeyIndex)

		rightNode = &node{
			leaf:  treeNode.leaf,
			keys:  []keyStruct{},
			child: []*node{},
		}

		for middleKeyIndex < len(treeNode.keys) {
			rightNode.keys = append(rightNode.keys, treeNode.keys[middleKeyIndex])
			treeNode.keys = popFromSlice(treeNode.keys, middleKeyIndex)
		}

		for middleKeyIndex+1 < len(treeNode.child) {
			rightNode.child = append(rightNode.child, treeNode.child[middleKeyIndex+1])
			treeNode.child = popFromSlice(treeNode.child, middleKeyIndex+1)
		}

		return true, middleKey, rightNode
	}
}

func has(slice []keyStruct, key string) bool {
	for _, i := range slice {
		if key == i.key {
			return true
		}
	}
	return false
}

// func removeItemFromSlice[T comparable](slice []T, key T) []T {
// 	i := 0
// 	for slice[i] != key {
// 		i++
// 		if i >= len(slice) {
// 			return slice
// 		}
// 	}
// 	return popFromSlice[T](slice, i)
// }

func removeKeyFromSlice(slice []keyStruct, key string) []keyStruct {
	i := 0
	for slice[i].key != key {
		i++
		if i >= len(slice) {
			return slice
		}
	}
	return popFromSlice(slice, i)
}

func (tree *Btree) Delete(key string) {
	balancingRequired := false
	i := 0
	if !has(tree.root.keys, key) {
		if tree.root.leaf {
			return // handle error future problem
		}
		i = 0
		for i < len(tree.root.keys) && key > tree.root.keys[i].key {
			i++
		}
		balancingRequired = tree.deleteHelper(tree.root.child[i], key)
	} else if tree.root.leaf {
		tree.root.keys = removeKeyFromSlice(tree.root.keys, key)
	} else {
		i = 0
		for i < len(tree.root.child) && key != tree.root.keys[i].key {
			i++
		}

		// find inorder predessecor and check if has more than minimum number of keys
		inorderPredecessorNode := tree.root.child[i]
		for !inorderPredecessorNode.leaf {
			inorderPredecessorNode = inorderPredecessorNode.child[len(inorderPredecessorNode.child)-1]
		}

		if len(inorderPredecessorNode.keys) > tree.maxElements {
			tree.root.keys[i] = inorderPredecessorNode.keys[len(inorderPredecessorNode.child)-1]
			inorderPredecessorNode.keys = popFromSlice(inorderPredecessorNode.keys, len(inorderPredecessorNode.keys)-1)
		} else if i < len(tree.root.child)-1 {
			// find inorder successor
			inorderSuccessorNode := tree.root.child[i+1]
			for !inorderSuccessorNode.leaf {
				inorderSuccessorNode = inorderSuccessorNode.child[0]
			}

			if len(inorderSuccessorNode.keys) > tree.minElements {
				tree.root.keys[i] = inorderSuccessorNode.keys[0]
				inorderSuccessorNode.keys = popFromSlice(inorderSuccessorNode.keys, 0)
			} else {
				tree.root.keys[i] = inorderSuccessorNode.keys[0]
				balancingRequired = tree.deleteHelper(tree.root.child[i+1], inorderSuccessorNode.keys[0].key)
			}
			i++ // changed where it was placed check here for likely bug
		} else {
			tree.root.keys[i] = inorderPredecessorNode.keys[len(inorderPredecessorNode.child)-1]
			balancingRequired = tree.deleteHelper(tree.root.child[i], inorderPredecessorNode.keys[len(inorderPredecessorNode.child)-1].key)
		}
	}
	// root handling of breaking tree conditions
	if balancingRequired {
		// check left sibling
		if i > 0 && len(tree.root.child[i-1].keys) > tree.minElements {
			// rotating keys around to balance tree
			tree.root.child[i].keys = insertIntoSlice(tree.root.child[i].keys, 0, tree.root.keys[i-1])
			tree.root.keys[i-1] = tree.root.child[i-1].keys[len(tree.root.child[i-1].keys)-1]
			tree.root.child[i-1].keys = popFromSlice(tree.root.child[i-1].keys, len(tree.root.child[i-1].keys)-1)

			if !tree.root.child[i].leaf {
				// handling child node
				tree.root.child[i].child = insertIntoSlice(tree.root.child[i].child, 0, tree.root.child[i-1].child[len(tree.root.child[i-1].child)-1])
				tree.root.child[i-1].child = popFromSlice(tree.root.child[i-1].child, len(tree.root.child[i-1].child)-1)
			}

		} else if i < len(tree.root.keys) && len(tree.root.child[i+1].keys) > tree.minElements {
			// rotating keys around to balance tree
			tree.root.child[i].keys = append(tree.root.child[i].keys, tree.root.keys[i])
			tree.root.keys[i] = tree.root.child[i+1].keys[0]
			tree.root.child[i+1].keys = popFromSlice(tree.root.child[i+1].keys, 0)

			if !tree.root.child[i].leaf {
				// handling child node
				tree.root.child[i].child = append(tree.root.child[i].child, tree.root.child[i+1].child[0])
				tree.root.child[i+1].child = popFromSlice(tree.root.child[i+1].child, 0)
			}

		} else {
			if i > 0 {
				tree.root.child[i-1].keys = append(tree.root.child[i-1].keys, tree.root.keys[i-1])
				tree.root.keys = popFromSlice(tree.root.keys, i-1)
				tree.root.child[i-1].keys = append(tree.root.child[i-1].keys, tree.root.child[i].keys...)

				if !tree.root.child[i].leaf {
					// handling children
					tree.root.child[i-1].child = append(tree.root.child[i-1].child, tree.root.child[i].child...)
				}

				tree.root.child = popFromSlice(tree.root.child, i)
			} else {
				tree.root.child[i].keys = append(tree.root.child[i].keys, tree.root.keys[i])
				tree.root.keys = popFromSlice(tree.root.keys, i)
				tree.root.child[i].keys = append(tree.root.child[i].keys, tree.root.child[i+1].keys...)

				if !tree.root.child[i].leaf {
					// handling children
					tree.root.child[i].child = append(tree.root.child[i].child, tree.root.child[i+1].child...)
				}

				tree.root.child = popFromSlice(tree.root.child, i+1)
			}

			if len(tree.root.keys) <= 0 {
				tree.root = tree.root.child[0]
			}
		}

	}
}

func (tree *Btree) deleteHelper(treeNode *node, key string) bool {
	balancingRequired := false
	i := 0
	if !has(treeNode.keys, key) {
		if treeNode.leaf {
			return false // handle error future problem
		}
		i = 0
		for i < len(treeNode.keys) && key > treeNode.keys[i].key {
			i++
		}
		balancingRequired = tree.deleteHelper(treeNode.child[i], key)
	} else if treeNode.leaf {
		treeNode.keys = removeKeyFromSlice(treeNode.keys, key)

		// check if node is breaking the minElements rule
		return len(treeNode.keys) < tree.minElements
	} else {
		i = 0
		for i < len(treeNode.child) && key != treeNode.keys[i].key {
			i++
		}

		// find inorder predessecor and check if has more than minimum number of keys
		inorderPredecessorNode := treeNode.child[i]
		for !inorderPredecessorNode.leaf {
			inorderPredecessorNode = inorderPredecessorNode.child[len(inorderPredecessorNode.child)-1]
		}

		if len(inorderPredecessorNode.keys) > tree.maxElements {
			treeNode.keys[i] = inorderPredecessorNode.keys[len(inorderPredecessorNode.child)-1]
			inorderPredecessorNode.keys = popFromSlice(inorderPredecessorNode.keys, len(inorderPredecessorNode.keys)-1)
		} else if i < len(treeNode.child)-1 {
			// find inorder successor
			inorderSuccessorNode := treeNode.child[i+1]
			for !inorderSuccessorNode.leaf {
				inorderSuccessorNode = inorderSuccessorNode.child[0]
			}

			if len(inorderSuccessorNode.keys) > tree.minElements {
				treeNode.keys[i] = inorderSuccessorNode.keys[0]
				inorderSuccessorNode.keys = popFromSlice(inorderSuccessorNode.keys, 0)
			} else {
				treeNode.keys[i] = inorderSuccessorNode.keys[0]
				balancingRequired = tree.deleteHelper(treeNode.child[i+1], inorderSuccessorNode.keys[0].key)
			}
			i++ // changed where it was placed check here for likely bug
		} else {
			treeNode.keys[i] = inorderPredecessorNode.keys[len(inorderPredecessorNode.child)-1]
			balancingRequired = tree.deleteHelper(treeNode.child[i], inorderPredecessorNode.keys[len(inorderPredecessorNode.child)-1].key)
		}
	}
	// root handling of breaking tree conditions
	if balancingRequired {
		// check left sibling
		if i > 0 && len(treeNode.child[i-1].keys) > tree.minElements {
			// rotating keys around to balance tree
			treeNode.child[i].keys = insertIntoSlice(treeNode.child[i].keys, 0, treeNode.keys[i-1])
			treeNode.keys[i-1] = treeNode.child[i-1].keys[len(treeNode.child[i-1].keys)-1]
			treeNode.child[i-1].keys = popFromSlice(treeNode.child[i-1].keys, len(treeNode.child[i-1].keys)-1)

			if !treeNode.child[i].leaf {
				// handling child treeNode
				treeNode.child[i].child = insertIntoSlice(treeNode.child[i].child, 0, treeNode.child[i-1].child[len(treeNode.child[i-1].child)-1])
				treeNode.child[i-1].child = popFromSlice(treeNode.child[i-1].child, len(treeNode.child[i-1].child)-1)
			}

		} else if i < len(treeNode.keys) && len(treeNode.child[i+1].keys) > tree.minElements {
			// rotating keys areound to balance tree
			treeNode.child[i].keys = append(treeNode.child[i].keys, treeNode.keys[i])
			treeNode.keys[i] = treeNode.child[i+1].keys[0]
			treeNode.child[i+1].keys = popFromSlice(treeNode.child[i+1].keys, 0)

			if !treeNode.child[i].leaf {
				// handling child treeNode
				treeNode.child[i].child = append(treeNode.child[i].child, treeNode.child[i+1].child[0])
				treeNode.child[i+1].child = popFromSlice(treeNode.child[i+1].child, 0)
			}

		} else {
			if i > 0 {
				treeNode.child[i-1].keys = append(treeNode.child[i-1].keys, treeNode.keys[i-1])
				treeNode.keys = popFromSlice(treeNode.keys, i-1)
				treeNode.child[i-1].keys = append(treeNode.child[i-1].keys, treeNode.child[i].keys...)

				if !treeNode.child[i].leaf {
					// handling children
					treeNode.child[i-1].child = append(treeNode.child[i-1].child, treeNode.child[i].child...)
				}
				treeNode.child = popFromSlice(treeNode.child, i)
			} else {
				treeNode.child[i].keys = append(treeNode.child[i].keys, treeNode.keys[i])
				treeNode.keys = popFromSlice(treeNode.keys, i)
				treeNode.child[i].keys = append(treeNode.child[i].keys, treeNode.child[i+1].keys...)

				if !treeNode.child[i].leaf {
					// handling children
					treeNode.child[i].child = append(treeNode.child[i].child, treeNode.child[i+1].child...)
				}
				treeNode.child = popFromSlice(treeNode.child, i+1)
			}

			if len(treeNode.keys) < tree.minElements {
				return true
			}

		}
	}
	return false
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
