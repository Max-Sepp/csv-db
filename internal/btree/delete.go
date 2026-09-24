package btree

import "errors"

var ErrKeyNotFound = errors.New("key not found in btree")

// Delete removes an entry with the given key and returns its row pointer
func (tree *Btree) Delete(key string) (int64, error) {
	return tree.deleteEntry(key, nil)
}

// DeleteEntry removes the entry with the given key and row pointer.
// This is needed when keys are not unique so the correct entry is removed.
func (tree *Btree) DeleteEntry(key string, rowPtr int64) error {
	_, err := tree.deleteEntry(key, &rowPtr)
	return err
}

func (tree *Btree) deleteEntry(key string, rowPtr *int64) (int64, error) {
	deletedRowPtr, err := deleteHelper(tree, nil, tree.root, key, rowPtr)

	if err != nil {
		return deletedRowPtr, err
	}

	for len(tree.root.keys) == 0 && len(tree.root.child) != 0 {
		tree.root = tree.root.child[0]
	}

	return deletedRowPtr, err
}

// deleteHelper removes an entry matching key (and targetRowPtr if it is not nil) from the subtree at currentNode.
// If no entry matches the subtree is left unchanged and ErrKeyNotFound is returned.
func deleteHelper(tree *Btree, parentNode *node, currentNode *node, key string, targetRowPtr *int64) (int64, error) {
	var rowPtr int64 = -1
	var indexOfChildNode int
	var err error

	if parentNode != nil {
		indexOfChildNode = parentNode.indexOfChildNode(currentNode)
	}

	matches := func(k keyStruct) bool {
		return k.key == key && (targetRowPtr == nil || k.rowPtr == *targetRowPtr)
	}

	// Duplicate keys can be spread over several keys of this node and the children between them,
	// so every position from the first key >= key up to the last key == key has to be checked.
	keyIndex := currentNode.findKeyIndex(key)
	found := false
	for {
		if keyIndex < len(currentNode.keys) && matches(currentNode.keys[keyIndex]) {
			found = true
			break
		}

		if !currentNode.leaf {
			rowPtr, err = deleteHelper(tree, currentNode, currentNode.child[keyIndex], key, targetRowPtr)
			if err == nil {
				break
			}
			if err != ErrKeyNotFound {
				return -1, err
			}
		}

		if keyIndex >= len(currentNode.keys) || currentNode.keys[keyIndex].key != key {
			return -1, ErrKeyNotFound
		}
		keyIndex++
	}

	if found && currentNode.leaf {
		rowPtr = currentNode.keys[keyIndex].rowPtr
		currentNode.keys = popFromSlice(currentNode.keys, keyIndex)
	} else if found {
		rowPtr = currentNode.keys[keyIndex].rowPtr

		inorderSuccessor := currentNode.getInorderSuccessor(keyIndex)

		currentNode.keys[keyIndex] = inorderSuccessor

		// remove exactly the successor entry, not just any entry with the same key
		_, err = deleteHelper(tree, currentNode, currentNode.child[keyIndex+1], inorderSuccessor.key, &inorderSuccessor.rowPtr)

		if err != nil {
			return -1, err
		}
	}

	if tree.violatesMinimumNumberKeys(currentNode) && parentNode != nil {
		if indexOfChildNode > 0 && tree.nodeCanBeBorrowedFrom(parentNode.child[indexOfChildNode-1]) {
			// check if immediate left sibling can be borrowed from
			rotateFromLeft(parentNode, indexOfChildNode)
		} else if indexOfChildNode+1 < len(parentNode.child) && tree.nodeCanBeBorrowedFrom(parentNode.child[indexOfChildNode+1]) {
			// check if immediate right sibling can be borrowed from
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
