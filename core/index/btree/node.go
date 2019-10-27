package btree

type Node struct {
	DataList []Data
	Parent *Node
	Depth int
}

func NewSingleDataNode(key interface{}, value interface{}) *Node{
	node := Node{}
	data := NewDataWithNode(key, value, &node)
	node.DataList = []Data{
		data,
	}
	return &node
}