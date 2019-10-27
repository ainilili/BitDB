package btree

type BTree struct {
	Root *Node
	Partition int
}

func (btree BTree) Insert(key interface{}, value interface{}) (interface{}, error){
	if btree.Root == nil{
		btree.Root = NewSingleDataNode(key, value)
	}
}