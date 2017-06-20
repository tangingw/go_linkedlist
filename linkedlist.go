package main

import (
	"fmt"
	"os"
)

/**
This is the go implementation of linked list. The code might be quite C-like
as it is borrowed from the examples in http://www.geeksforgeeks.org/linked-list-set-1-introduction/
**/

type Node struct {
	Next  *Node
	Value interface{}
}

func (n *Node) Read() {
	fmt.Println(n.Value)
}

func (n *Node) Write(value interface{}) {
	n.Value = value
}

func (n *Node) Append(value interface{}) {

	newNode := new(Node)
	lastNode := n

	newNode.Value = value
	newNode.Next = nil

	if n == nil {
		n = newNode

	}

	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}

	lastNode.Next = newNode

}

func (n *Node) Push(newValue interface{}) *Node {

	newNode := new(Node)
	newNode.Value = newValue
	newNode.Next = n

	return newNode
}

func (n *Node) InsertAfter(newValue interface{}) {

	if n == nil {
		fmt.Println("The given previous node cannot be nil")
		os.Exit(1)
	}

	newNode := new(Node)
	newNode.Value = newValue
	newNode.Next = n.Next

	n.Next = newNode
}

func (n *Node) Right() *Node {

	return n.Next
}

func (n *Node) Left(value interface{}) *Node {

	previousNode := new(Node)

	previousNode = n

	for n != nil {

		if n.Value == value {

			if previousNode.Value == value {
				fmt.Println("This is the edge of the linkedlist")
			}
			return previousNode
		}

		previousNode = n
		n = n.Next
	}

	return nil
}

func (n *Node) DeleteNode(key interface{}) {

	temp := n
	previousNode := new(Node)

	if temp != nil && temp.Value == key {
		n = temp.Next
		return
	}

	for temp != nil && temp.Value != key {

		previousNode = temp
		temp = temp.Next
	}

	if temp == nil {
		return
	}

	previousNode.Next = temp.Next
}

func (n *Node) Search(key interface{}) bool {

	for n != nil {

		if n.Value == key {
			return true
		}

		n = n.Next
	}

	return false
}

func (n *Node) PrintList() {

	for n != nil {
		fmt.Println(n.Value)
		n = n.Next
	}
}

func (n *Node) Reverse() *Node {

	//Get the value from the first link
	newNode := new(Node)
	newNode.Value = n.Value

	n = n.Next

	//Loop the remaining links
	for n != nil {

		newNode = newNode.Push(n.Value)
		n = n.Next
	}

	return newNode
}
