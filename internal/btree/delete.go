package btree

func (tree *Btree) Delete(key string) (int64, error) {
	parentNode, nodeContainingKey := search(nil, tree.root, key)

	if nodeContainingKey.leaf {
		removeKeyFromSlice(nodeContainingKey.keys, key)

		indexOfChildNode := parentNode.indexOfChildNode(nodeContainingKey)

		if tree.violatesMinimumNumberKeys(nodeContainingKey) {
			// check if immediate left sibling can be borrowed from
			if indexOfChildNode-1 >= 0 && tree.nodeCanBeBorrowedFrom(parentNode.child[indexOfChildNode-1]) {
				// rotating keys around to balance tree
				// TODO: once completed Delete check if this code repeated and remove into own function if repeated
				parentNode.child[indexOfChildNode].keys = insertIntoSlice(parentNode.child[indexOfChildNode].keys, 0, parentNode.keys[indexOfChildNode-1])
				parentNode.keys[indexOfChildNode-1] = parentNode.child[indexOfChildNode-1].keys[len(parentNode.child[indexOfChildNode-1].keys)-1]
				parentNode.child[indexOfChildNode-1].keys = popFromSlice(parentNode.child[indexOfChildNode-1].keys, len(parentNode.child[indexOfChildNode-1].keys)-1)

			} else if indexOfChildNode+1 < len(parentNode.child) && tree.nodeCanBeBorrowedFrom(parentNode.child[indexOfChildNode+1]) {
				// rotating keys areound to balance tree
				parentNode.child[indexOfChildNode].keys = append(parentNode.child[indexOfChildNode].keys, parentNode.keys[indexOfChildNode])
				parentNode.keys[indexOfChildNode] = parentNode.child[indexOfChildNode+1].keys[0]
				parentNode.child[indexOfChildNode+1].keys = popFromSlice(parentNode.child[indexOfChildNode+1].keys, 0)
			} else {

			}
		}
	}

}

func search(parentNode *node, treeNode *node, key string) (*node, *node) {
	if treeNode == nil {
		return nil, nil
	}

	// This is either the location of the treeNode or the index of the child which should be checked next to find the treeNode
	keyIndex := treeNode.findKeyIndex(key)

	if treeNode.keys[keyIndex].key == key {
		return parentNode, treeNode
	}

	return search(treeNode, treeNode.child[keyIndex], key)
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
