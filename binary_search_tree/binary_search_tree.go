package binary_search_tree

import "math"

type Node struct {
	left  *Node
	right *Node
	key   int
}

func Insert(head *Node, key int) *Node {
	if head == nil {
		newNode := Node{
			left:  nil,
			right: nil,
			key:   key,
		}
		return &newNode
	}

	if key < head.key {
		*head = Node{
			left:  Insert(head.left, key),
			right: head.right,
			key:   head.key,
		}
	}

	if key > head.key {
		*head = Node{
			left:  head.left,
			right: Insert(head.right, key),
			key:   head.key,
		}
	}

	return head
}

func DFS(head *Node, target int) *Node {
	if head == nil {
		return nil
	}

	if head.key < target {
		return DFS(head.right, target)
	}

	if head.key > target {
		return DFS(head.left, target)
	}

	return head
}

func BFS(head *Node, target int) *Node {
	if head == nil {
		return nil
	}

	if head.key == target {
		return head
	}
	if head.key > target {
		return BFS(head.left, target)
	}
	return BFS(head.right, target)
}

func Height(head *Node) int {
	if head == nil {
		return 0
	}

	return int(math.Max(float64(Height(head.left)), float64(Height(head.right)))) + 1
}

func Size(head *Node) int {
	if head == nil {
		return 0
	}

	return Size(head.left) + Size(head.right) + 1
}
