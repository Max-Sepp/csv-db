package btree

func (tree *Btree) Delete(key string) (int64, error) {
	rowPtr, err := deleteHelper(tree, nil, tree.root, key)

	if err != nil {
		return rowPtr, err
	}

	for len(tree.root.keys) == 0 && len(tree.root.child) != 0 {
		tree.root = tree.root.child[0]
	}

	return rowPtr, err
}

func deleteHelper(tree *Btree, parentNode *node, currentNode *node, key string) (int64, error) {
	var rowPtr int64 = -1
	var indexOfChildNode int
	var err error

	// This is either the location of the treeNode or the index of the child which should be checked next to find the treeNode
	keyIndex := currentNode.findKeyIndex(key)

	if parentNode != nil {
		indexOfChildNode = parentNode.indexOfChildNode(currentNode)
	}

	if keyIndex >= len(currentNode.keys) || currentNode.keys[keyIndex].key != key {
		rowPtr, err = deleteHelper(tree, currentNode, currentNode.child[keyIndex], key)

		if err != nil {
			return -1, err
		}

	} else if currentNode.leaf {
		currentNode.keys, rowPtr = removeKeyFromSlice(currentNode.keys, key)
	} else {
		inorderSuccessor := currentNode.getInorderSuccessor(keyIndex)

		currentNode.keys[keyIndex] = inorderSuccessor

		rowPtr, err = deleteHelper(tree, currentNode, currentNode.child[keyIndex+1], inorderSuccessor.key)

		if err != nil {
			return -1, err
		}
	}

	if tree.violatesMinimumNumberKeys(currentNode) && parentNode != nil {

		// check if immediate left sibling can be borrowed from
		if indexOfChildNode > 0 && tree.nodeCanBeBorrowedFrom(parentNode.child[indexOfChildNode-1]) {
			rotateFromLeft(parentNode, indexOfChildNode)

			// check if immediate right sibling can be borrowed from
		} else if indexOfChildNode+1 < len(parentNode.child) && tree.nodeCanBeBorrowedFrom(parentNode.child[indexOfChildNode+1]) {
			rotateFromRight(parentNode, indexOfChildNode)
		} else {
			handleProblemChild(parentNode, indexOfChildNode)
		}
	}

	return rowPtr, nil
}

func rotateFromLeft(parentNode *node, indexOfChildNode int) {
	parentNode.child[indexOfChildNode].keys = insertIntoSlice(parentNode.child[indexOfChildNode].keys, 0, parentNode.keys[indexOfChildNode-1])
	parentNode.keys[indexOfChildNode-1] = parentNode.child[indexOfChildNode-1].keys[len(parentNode.child[indexOfChildNode-1].keys)-1]
	parentNode.child[indexOfChildNode-1].keys = popFromSlice(parentNode.child[indexOfChildNode-1].keys, len(parentNode.child[indexOfChildNode-1].keys)-1)

	if !parentNode.child[indexOfChildNode].leaf {
		// handling child treeNode
		parentNode.child[indexOfChildNode].child = insertIntoSlice(parentNode.child[indexOfChildNode].child, 0, parentNode.child[indexOfChildNode-1].child[len(parentNode.child[indexOfChildNode-1].child)-1])
		parentNode.child[indexOfChildNode-1].child = popFromSlice(parentNode.child[indexOfChildNode-1].child, len(parentNode.child[indexOfChildNode-1].child)-1)
	}
}

func rotateFromRight(parentNode *node, indexOfChildNode int) {
	parentNode.child[indexOfChildNode].keys = append(parentNode.child[indexOfChildNode].keys, parentNode.keys[indexOfChildNode])
	parentNode.keys[indexOfChildNode] = parentNode.child[indexOfChildNode+1].keys[0]
	parentNode.child[indexOfChildNode+1].keys = popFromSlice(parentNode.child[indexOfChildNode+1].keys, 0)

	if !parentNode.child[indexOfChildNode].leaf {
		// handling child treeNode
		parentNode.child[indexOfChildNode].child = append(parentNode.child[indexOfChildNode].child, parentNode.child[indexOfChildNode+1].child[0])
		parentNode.child[indexOfChildNode+1].child = popFromSlice(parentNode.child[indexOfChildNode+1].child, 0)
	}
}

func handleProblemChild(parentNode *node, problemChildIndex int) {
	if problemChildIndex > 0 {
		parentNode.child[problemChildIndex-1].keys = append(parentNode.child[problemChildIndex-1].keys, parentNode.keys[problemChildIndex-1])
		parentNode.keys = popFromSlice(parentNode.keys, problemChildIndex-1)

		parentNode.child[problemChildIndex-1].keys = append(parentNode.child[problemChildIndex-1].keys, parentNode.child[problemChildIndex].keys...)

		if !parentNode.child[problemChildIndex].leaf {
			parentNode.child[problemChildIndex-1].child = append(parentNode.child[problemChildIndex-1].child, parentNode.child[problemChildIndex].child...)
		}

		parentNode.child = popFromSlice(parentNode.child, problemChildIndex)
	} else {
		parentNode.child[problemChildIndex].keys = append(parentNode.child[problemChildIndex].keys, parentNode.keys[problemChildIndex])
		parentNode.keys = popFromSlice(parentNode.keys, problemChildIndex)

		parentNode.child[problemChildIndex].keys = append(parentNode.child[problemChildIndex].keys, parentNode.child[problemChildIndex+1].keys...)

		if !parentNode.child[problemChildIndex+1].leaf {
			parentNode.child[problemChildIndex].child = append(parentNode.child[problemChildIndex].child, parentNode.child[problemChildIndex+1].child...)
		}

		parentNode.child = popFromSlice(parentNode.child, problemChildIndex+1)
	}
}

func removeKeyFromSlice(slice []keyStruct, key string) ([]keyStruct, int64) {
	i := 0
	for slice[i].key != key {
		i++
		if i >= len(slice) {
			return slice, -1
		}
	}

	offset := slice[i].rowPtr
	return popFromSlice(slice, i), offset
}

func (btree *Btree) violatesMinimumNumberKeys(treeNode *node) bool {
	return btree.minElements > len(treeNode.keys)
}

func (btree *Btree) nodeCanBeBorrowedFrom(treeNode *node) bool {
	return btree.minElements < len(treeNode.keys)
}

// assumes targetNode is actually a child of parentNode
func (parentNode *node) indexOfChildNode(targetNode *node) int {
	i := 0
	for i < len(parentNode.child) && parentNode.child[i] != targetNode {
		i++
	}
	return i
}

func (treeNode *node) getInorderSuccessor(keyIndex int) keyStruct {
	inorderSuccessorNode := treeNode.child[keyIndex+1]
	for !inorderSuccessorNode.leaf {
		inorderSuccessorNode = inorderSuccessorNode.child[0]
	}
	return inorderSuccessorNode.keys[0]
}
