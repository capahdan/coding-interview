package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	node := Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}

	Bfs(&node)
}

func Bfs(node *Node) {

	var arr []*Node
	arr = append(arr, node)

	for len(arr) > 0 {
		node = arr[0]
		arr = arr[1:]
		fmt.Println(node.Val)

		if node.Left != nil {
			arr = append(arr, node.Left)
		}
		if node.Right != nil {
			arr = append(arr, node.Right)
		}
	}

}
