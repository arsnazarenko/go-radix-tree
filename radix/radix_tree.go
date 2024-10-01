package radix

import (
	"fmt"
)

// Node represents a node in the Radix Tree.
type Node[T any] struct {
	Value      T
	Children   map[rune]*Node[T]
	IsTerminal bool
}

// RadixTree represents a Radix Tree.
type RadixTree[T any] struct {
	root *Node[T]
}

// NewRadixTree creates a new Radix Tree.
func NewRadixTree[T any]() *RadixTree[T] {
	return &RadixTree[T]{
		root: &Node[T]{
			Children: make(map[rune]*Node[T]),
		},
	}
}

// Insert inserts a new key-value pair into the Radix Tree.
func (rt *RadixTree[T]) Insert(key string, value T) {
	node := rt.root

	for _, r := range key {
		if _, ok := node.Children[r]; !ok {
			node.Children[r] = &Node[T]{
				Children: make(map[rune]*Node[T]),
			}
		}
		node = node.Children[r]
	}

	node.Value = value
	node.IsTerminal = true
}

// Search searches for a key in the Radix Tree and returns its associated value.
func (rt *RadixTree[T]) Search(key string) (T, bool) {
	node := rt.root
    var dummy T
	for _, r := range key {
		if _, ok := node.Children[r]; !ok {
			return dummy, false
		}
		node = node.Children[r]
	}

	if node.IsTerminal {
		return node.Value, true
	}

	return dummy, false
}

// Delete deletes a key from the Radix Tree.
func (rt *RadixTree[T]) Delete(key string) {
	node := rt.root
	path := make([]rune, 0)

	for _, r := range key {
		if _, ok := node.Children[r]; !ok {
			return
		}
		path = append(path, r)
		node = node.Children[r]
	}

	// If node is a terminal node, simply mark it as non-terminal.
	if node.IsTerminal {
		node.IsTerminal = false
		return
	}

	// If node has children, delete the key prefix up to the last shared prefix.
	for i := len(path) - 1; i >= 0; i-- {
		parent := rt.getNode(path[:i])
		if len(parent.Children) > 1 {
			return
		}
		delete(parent.Children, path[i])
	}
}

// getNode returns the node at the given path.
func (rt *RadixTree[T]) getNode(path []rune) *Node[T] {
	node := rt.root
	for _, r := range path {
		node = node.Children[r]
	}
	return node
}

// Print prints the Radix Tree in a human-readable format.
func (rt *RadixTree[T]) Print() {
	fmt.Println("Radix Tree:")
	rt.printNode(rt.root, "")
}

func (rt *RadixTree[T]) printNode(node *Node[T], prefix string) {
	for r, child := range node.Children {
		newPrefix := prefix + string(r)
		if child.IsTerminal {
			fmt.Printf("%s: %vn", newPrefix, child.Value)
		}
		rt.printNode(child, newPrefix)
	}
}
