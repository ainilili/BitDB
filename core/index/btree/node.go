package btree

type Node struct {
	DataList []Data
	Parent Node
}
