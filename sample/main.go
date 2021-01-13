package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	node := read()

	for _, value := range node {
		printNode(value)
		fmt.Println()
	}
}

func read() []Node {
	var N int

	fmt.Scanf("%d", &N)
	//fmt.Print("Value of N is ", N)
	var node = make([]Node, N)
	//fmt.Println()
	for i := 0; i < N; i++ {
		var Value, LeftNodeIndex, RightNodeIndex int
		fmt.Scanf("%d  %d %d", &Value, &LeftNodeIndex, &RightNodeIndex)
		//fmt.Println(Value, LeftNodeIndex, RightNodeIndex)

		node[i].Value = Value

		if LeftNodeIndex != -1 {
			node[i].Left = &node[LeftNodeIndex]
		}

		if RightNodeIndex != -1 {
			node[i].Right = &node[RightNodeIndex]
		}

	}

	return node

}

func printNode(n Node) {
	fmt.Print(n.Value, " ")
	if n.Left != nil {
		fmt.Print(n.Left.Value, " ")
	}

	if n.Right != nil {
		fmt.Println(n.Right.Value, " ")
	}
}
