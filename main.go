package main

import (
	"fmt"
)

func main() {

	head := new(Node)
	head.Value = -1

	head = head.Push(-3)
	head = head.Push(-2)

	head.Append(4)
	head.InsertAfter(3)
	head.Next.InsertAfter(6)
	//head.PrintList()

	//fmt.Println(head.Next.Next)

	head.DeleteNode(3)
	head.Next.Write(100)
	//head.PrintList()
	//fmt.Println(head.Search(-3))

	head.PrintList()
	//temp := head.Reverse()
	//temp.PrintList()
	/**
		It means find the link where it is on the left of the link with value 4
	**/
	fmt.Println(head.Left(4))
}
