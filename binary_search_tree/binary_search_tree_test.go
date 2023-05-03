package binary_search_tree

import (
	"testing"
)

func TestHead(t *testing.T) {
	var head *Node

	head = Insert(head, 1)
}

func TestInsert(t *testing.T) {
	var head *Node

	values := []int{5, 1, 3, 7}

	for _, value := range values {
		head = Insert(head, value)
	}

	if head.key != 5 || head.left.key != 1 || head.right.key != 7 {
		t.Fatalf(`Insert(head, value) = *Node should inert elements in order`)
	}
}

func TestDFSFound(t *testing.T) {
	var head *Node

	values := []int{5, 1, 3, 7}

	for _, value := range values {
		head = Insert(head, value)
	}

	for _, value := range values {
		resNode := DFS(head, value)
		if resNode == nil {
			t.Fatalf(`DFS(head, %v) = %v shoud not be nil`, value, resNode)
		}
	}
}

func TestDFSNotFound(t *testing.T) {
	var head *Node

	values := []int{5, 1, 3, 7}
	remainingValues := []int{2, 4, 6, 8}

	for _, value := range values {
		head = Insert(head, value)
	}

	for _, value := range remainingValues {
		resNode := DFS(head, value)
		if resNode != nil {
			t.Fatalf(`DFS(head, %v) = %v shoud be nil`, value, resNode)
		}
	}
}

func TestBFSFound(t *testing.T) {
	var head *Node

	values := []int{5, 1, 3, 7}

	for _, value := range values {
		head = Insert(head, value)
	}

	for _, value := range values {
		resNode := BFS(head, value)
		if resNode == nil {
			t.Fatalf(`BFS(head, %v) = %v shoud not be nil`, value, resNode)
		}
	}
}

func TestBFSNotFound(t *testing.T) {
	var head *Node

	values := []int{5, 1, 3, 7}
	remainingValues := []int{2, 4, 6, 8}

	for _, value := range values {
		head = Insert(head, value)
	}

	for _, value := range remainingValues {
		resNode := BFS(head, value)
		if resNode != nil {
			t.Fatalf(`BFS(&head, %v) = %v shoud be nil`, value, resNode)
		}
	}
}

func TestHeight(t *testing.T) {
	var head *Node
	var h int
	h = Height(head)

	if h != 0 {
		t.Fatalf(`Height(head) = %v shoud return 0`, h)
	}

	values := []int{5, 1, 3, 7}

	for _, value := range values {
		head = Insert(head, value)
	}

	h = Height(head)

	if h != 3 {
		t.Fatalf(`Height(head) = %v shoud return 3`, h)
	}
}

func TestSize(t *testing.T) {
	var head *Node
	var s int
	s = Size(head)

	if s != 0 {
		t.Fatalf(`Size(head) = %v shoud return 0`, s)
	}

	values := []int{5, 1, 3, 7}

	for _, value := range values {
		head = Insert(head, value)
	}

	s = Size(head)

	if s != 4 {
		t.Fatalf(`Size(head) = %v shoud return 4`, s)
	}
}
