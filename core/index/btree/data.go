package btree

type Data struct {
	Key interface{}
	Value interface{}
	Parent interface{}
	LeftNode *Data
	RightNode *Data
	Depth int
}