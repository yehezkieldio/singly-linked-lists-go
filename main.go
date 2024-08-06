package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtBack(data int) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (list *LinkedList) insertAtFront(data int) {
	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.head = newNode
		return
	}

	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

func (list *LinkedList) insertAfterValue(afterValue, data int) {
	newNode := &Node{data: data, next: nil}

	current := list.head

	for current != nil {
		if current.data == afterValue {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}

	fmt.Printf("Value %d not found in the list\n", afterValue)
}

func (list *LinkedList) insertBeforeValue(beforeValue, data int) {
	if list.head == nil {
		return
	}

	if list.head.data == beforeValue {
		newNode := &Node{data: data, next: list.head}
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		if current.next.data == beforeValue {
			newNode := &Node{data: data, next: current.next}
			current.next = newNode
			return
		}
		current = current.next
	}

	fmt.Printf("Value %d not found in the list\n", beforeValue)
}

func (list *LinkedList) deleteFront() {
	if list.head != nil {
		list.head = list.head.next
		fmt.Printf("Deleted front node\n")
		return
	}
}

func (list *LinkedList) deleteLast() {
	if list.head == nil {
		fmt.Printf("List is empty\n")
		return
	}

	if list.head.next == nil {
		list.head = nil
		fmt.Printf("Deleted last node\n")
		return
	}

	var current *Node = list.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil

	fmt.Printf("Deleted last node\n")
}

func (list *LinkedList) print() {
	var current *Node = list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}

	list.insertAtBack(1)
	list.insertAtBack(2)

	list.deleteLast()
	list.insertAtBack(3)

	list.insertAtFront(0)

	list.insertAfterValue(2, 4)

	list.insertBeforeValue(4, 5)

	list.insertAtBack(100)

	list.deleteFront()

	list.print()
}
