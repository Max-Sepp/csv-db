package btree

// bool value says if item found or not
func (tree *Btree) Find(keyStr string) (bool, string, int64) {
	return tree.findHelper(tree.root, keyStr)
}

func (tree *Btree) findHelper(treeNode *node, keyStr string) (bool, string, int64) {
	i := 0
	for i < len(treeNode.keys) && keyStr > treeNode.keys[i].key {
		i++
	}

	if keyStr == treeNode.keys[i].key {
		return true, treeNode.keys[i].key, treeNode.keys[i].rowPtr
	} else if treeNode.leaf {
		return false, "", 0
	} else {
		return tree.findHelper(treeNode.child[i], keyStr)
	}
}
