package btree

type Data struct {
	Key interface{}
	Value interface{}
	SelfNode *Node
	LeftNode *Node
	RightNode *Node
}

func NewData(key interface{}, value interface{}) Data{
	return NewDataWithNode(key, value, nil)
}

func NewDataWithNode(key interface{}, value interface{}, selfNode *Node) Data{
	return Data{Key: key, Value: value, SelfNode: selfNode}
}
