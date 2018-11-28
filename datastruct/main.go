package main

import (
	"fmt"
)

type Node struct {
	Data string
	Next *Node
}
                                                
type List struct {
	HeadNode *Node
	TailNode *Node
}

func (this *List) InsetList(v string) {
	node := &Node{Data: v}
	this.TailNode.Next = node

	this.TailNode = node
}

func main() {
	fmt.Println("算法")
}
