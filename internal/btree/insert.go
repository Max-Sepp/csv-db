package btree

func (tree *Btree) Insert(keyStr string, rowPtr int64) {
	key := keyStruct{
		key:    keyStr,
		rowPtr: rowPtr,
	}
	rebalanceRequired, middleKey, rightNode := tree.insertHelper(tree.root, key)

	if !rebalanceRequired {
		return
	}

	// rebalance root node
	leftNodeTemp := tree.root
	tree.root = &node{
		leaf:  false,
		keys:  []keyStruct{middleKey},
		child: []*node{leftNodeTemp, rightNode},
	}
}

func (tree *Btree) insertHelper(treeNode *node, key keyStruct) (rebalanceRequired bool, middleKey keyStruct, rightNode *node) {
	if treeNode.leaf {
		// add key to keys of treeNode since it is a leaf treeNode and has space
		addKeyToTreenode(key, treeNode)

		if !tree.isNodeOverMaxKeyLimit(treeNode) {
			return false, keyStruct{}, nil
		}

		return tree.splitTreenode(treeNode)
	} else {
		placeToInsert := treeNode.findPlaceToInsertKey(key)

		rebalanceRequired, middleKey, rightNode := tree.insertHelper(treeNode.child[placeToInsert], key)

		if !rebalanceRequired {
			return false, keyStruct{}, nil
		}

		// add returned key and right node
		placeToInsert = treeNode.findPlaceToInsertKey(middleKey)
		treeNode.keys = insertIntoSlice(treeNode.keys, placeToInsert, middleKey)
		treeNode.child = insertIntoSlice(treeNode.child, placeToInsert+1, rightNode)

		if !tree.isNodeOverMaxKeyLimit(treeNode) {
			return false, keyStruct{}, nil
		}

		// split treeNode
		return tree.splitTreenode(treeNode)
	}
}

func addKeyToTreenode(key keyStruct, treeNode *node) {
	i := 0
	for i < len(treeNode.keys) && key.key > treeNode.keys[i].key {
		i++
	}
	treeNode.keys = insertIntoSlice(treeNode.keys, i, key)
}

func (tree *Btree) isNodeOverMaxKeyLimit(treeNode *node) bool {
	return len(treeNode.keys) > tree.maxElements
}

func (tree *Btree) splitTreenode(treeNode *node) (rebalanceRequired bool, middleKey keyStruct, rightNode *node) {
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

	for middleKeyIndex < len(treeNode.child)-1 {
		rightNode.child = append(rightNode.child, treeNode.child[middleKeyIndex+1])
		treeNode.child = popFromSlice(treeNode.child, middleKeyIndex+1)
	}

	return true, middleKey, rightNode
}

func (treeNode *node) findPlaceToInsertKey(key keyStruct) int {
	i := 0
	for i < len(treeNode.keys) && key.key > treeNode.keys[i].key {
		i++
	}
	return i
}
