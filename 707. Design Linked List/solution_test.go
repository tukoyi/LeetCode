package myList

import (
	"fmt"
	"testing"
)

func TestMyList(t *testing.T) {
	newLst := Constructor()
	newList := &newLst
	fmt.Printf("%s\n", newList)

	newList.AddAtHead(1)
	fmt.Printf("%s\n", newList)

	newList.AddAtTail(3)
	fmt.Printf("%s\n", newList)

	newList.AddAtIndex(1, 2)
	fmt.Printf("%s\n", newList)

	fmt.Printf("%d\n", newList.Get(1))

	newList.DeleteAtIndex(1)
	fmt.Printf("%s\n", newList)

	fmt.Printf("%d\n", newList.Get(1))

}
