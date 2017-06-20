# go Linked List
This 

Usage:
```
package main

import (
	"fmt"
)

func main() {

	head := new(Node) //Define a new linked list
	head.Value = -1 //Add a value to the first node of linked list

	head = head.Push(-3) //Push a new node with value -3 to the front of head
	head = head.Push(-2) //Push a new node with value -2 to the front of head

	head.Append(4) //Append a node to a linked lust
	head.InsertAfter(3) //Add a node after the node with value 3
	head.Next.InsertAfter(6) 
	head.PrintList() //Print the whole linked list

	head.DeleteNode(3) //Delete the Node with value 3 in the linked lust
	head.Next.Write(100) //Write the value
	//head.PrintList()
	//fmt.Println(head.Search(-3))

	head.PrintList()
	temp := head.Reverse() //Reverse the linked list
	temp.PrintList()
	/**
		It means find the link where it is on the left of the link with value 4
	**/
	fmt.Println(head.Left(4)) 
}

```
Reference: http://www.geeksforgeeks.org/linked-list-set-1-introduction/