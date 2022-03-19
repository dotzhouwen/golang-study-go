package main

type BinTreeNode struct {
	data interface{}
	parent *BinTreeNode
	lChild *BinTreeNode
	rChild *BinTreeNode
	height int
	size int
}

func NewBinTreeNode(e interface{}) *BinTreeNode {
	return &BinTreeNode{data: e, size: 1}
}

func (b *BinTreeNode) GetData() interface{} {
	if b == nil {
		return nil
	}
	return b.data
}

func (b *BinTreeNode) SetData(e interface{}) {
	b.data = e
}

func (b *BinTreeNode) HasParent() bool {
	return b.parent != nil
}

func (b *BinTreeNode) GetParent() *BinTreeNode {
	if !b.HasParent() {
		return nil
	}
	return b.parent
}

func (b *BinTreeNode) setParent(p *BinTreeNode) {
	b.parent = p
}





func main() {

}
