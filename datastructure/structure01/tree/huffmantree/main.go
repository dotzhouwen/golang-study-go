package main

import (
	"fmt"
	"sort"
)

type HuffmanTreeNode struct {
	Data   string
	Weight int
	Left   *HuffmanTreeNode
	Right  *HuffmanTreeNode
	Parent *HuffmanTreeNode
}

func NewHuffmanTreeNode(data string, weight int) *HuffmanTreeNode {
	return &HuffmanTreeNode{
		Data:   data,
		Weight: weight,
	}
}

func (n *HuffmanTreeNode) String() string {
	return fmt.Sprintf("{data:%s weight:%d}", n.Data, n.Weight)
}

type HuffmanTreeNodeList []*HuffmanTreeNode

func (h HuffmanTreeNodeList) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight
}

func (h HuffmanTreeNodeList) Len() int {
	return len(h)
}

func (h HuffmanTreeNodeList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

type HuffmanTree struct {
	Leaf   HuffmanTreeNodeList
	Code   map[string]string
	UnCode map[string]string
	Root   *HuffmanTreeNode
	Src    map[string]int
}

func (h *HuffmanTree) Init(src map[string]int) {
	h.Src = src
	h.Leaf = make(HuffmanTreeNodeList, len(src))
	h.Code = make(map[string]string, len(src))

	var i int
	for v, w := range h.Src {
		node := &HuffmanTreeNode{
			Data:   v,
			Weight: w,
		}
		h.Leaf[i] = node
		i++
	}
	sort.Sort(h.Leaf)
}

func (h *HuffmanTree) Build() {
	if h.Leaf.Len() < 1 {
		panic("error")
	}

	nodeList := h.Leaf
	for len(nodeList) > 1 {
		leftNode := nodeList[0]
		rightNode := nodeList[1]

		tempNode := &HuffmanTreeNode{
			Weight: leftNode.Weight + rightNode.Weight,
			Left:   leftNode,
			Right:  rightNode,
		}

		leftNode.Parent = tempNode
		rightNode.Parent = tempNode

		nodeList = regroup(nodeList[2:], tempNode)
	}
	h.Root = nodeList[0]
}

func (h *HuffmanTree) GetCode() {
	nodeList := h.Leaf
	for _, node := range nodeList {
		var s string
		curNode := node
		for curNode.Parent != nil {
			if curNode.Parent.Left == curNode {
				s = "0" + s
			} else {
				s = "1" + s
			}
			curNode = curNode.Parent
		}
		h.Code[node.Data] = s
	}

	h.UnCode = make(map[string]string, len(h.Code))
	for k, v := range h.Code {
		h.UnCode[v] = k
	}
}

func regroup(src HuffmanTreeNodeList, tempNode *HuffmanTreeNode) HuffmanTreeNodeList {
	size := len(src)
	result := make(HuffmanTreeNodeList, size+1)

	if size < 1 {
		result[0] = tempNode
		return result
	}

	if tempNode.Weight >= src[size-1].Weight {
		copy(result[:size], src[:])
		result[size] = tempNode
		return result
	}

	for i := 0; i < size; i++ {
		if src[i].Weight < tempNode.Weight {
			result[i] = src[i]
		} else {
			result[i] = tempNode
			copy(result[i+1:], src[i:])
			break
		}
	}
	return result
}

func (h *HuffmanTree) Encoding(s string) string {
	codeMap := h.Code
	var encodingStr string
	for _, v := range s {
		if c, ok := codeMap[string(v)]; ok {
			encodingStr += c
		}
	}
	return encodingStr
}

func (h *HuffmanTree) Decoding(encodeStr string) string {
	var decodeString string
	var basicUnit string

	for _, b := range encodeStr {
		basicUnit += string(b)
		if v, ok := h.UnCode[basicUnit]; ok {
			decodeString += v
			basicUnit = ""
		}
	}
	return decodeString
}

func main() {
	src := map[string]int{
		"A": 9,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 5,
		"F": 6,
	}

	huffmanTree := new(HuffmanTree)
	huffmanTree.Init(src)
	huffmanTree.Build()
	huffmanTree.GetCode()

	fmt.Println(huffmanTree.Root)

	fmt.Println("code:")
	fmt.Printf("%#v\n", huffmanTree.Code)

	srcStr := "ABCABCDE"
	encodeStr := huffmanTree.Encoding(srcStr)
	fmt.Println("encoding:", encodeStr)

	decodeStr := huffmanTree.Decoding(encodeStr)
	fmt.Println("decoding:", decodeStr)
}
