package linkedlist

import (
	"errors"
	"runtime"
)

type LinkedList interface {
	Add(*Node, string, string) error
	Get(*Node, string) (*Node, error)
	Delete(*Node, string) error
}

type Node struct {
	Key      string
	Value    string
	NextNode *Node
}

func NewLinkedList() *Node {
	return &Node{}
}

func Add(node *Node, key string, value string) error {
	nextNode := node
	for nextNode.NextNode != nil {
		nextNode = nextNode.NextNode
	}
	nextNode.NextNode = &Node{
		Key:      key,
		Value:    value,
		NextNode: nil}
	return nil
}

func Get(node *Node, key string) (*Node, error) {
	nextNode := node
	for nextNode.Key != key || nextNode.NextNode != nil {
		nextNode = nextNode.NextNode
	}
	return nextNode, nil
}

func Delete(node *Node, key string) error {
	var nextNode *Node = node.NextNode
	var currentNode *Node = node
	for nextNode.Key != key {
		currentNode = nextNode
		nextNode = nextNode.NextNode
		if currentNode.NextNode == nil {
			return errors.New("No node")
		}
	}
	currentNode.NextNode = nextNode.NextNode
	runtime.GC()
	return nil
}
