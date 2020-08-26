package main

import "fmt"

type Node struct {
	value      int
	next, prev *Node
}

type DoubleLinkedList struct {
	head, tail *Node
}

func (m *DoubleLinkedList) First() *Node {
	return m.head
}

func (m *Node) Next() *Node {
	return m.next
}

func (m *Node) Prev() *Node {
	return m.prev
}

func (m *DoubleLinkedList) Push(value int) {
	n := &Node{value: value}

	if m.head != nil {
		m.tail.next = n
		n.prev = m.tail
	} else {
		m.head = n
	}

	m.tail = n
}

func (m *DoubleLinkedList) Find(value int) *Node {
	foundFlag := false
	var foundNode *Node = nil

	for n := m.First(); n != nil && !foundFlag; n = n.Next() {
		if n.value == value {
			foundFlag = true
			foundNode = n
		}
	}

	return foundNode
}

func (m *DoubleLinkedList) Delete(value int) bool {
	node2del := m.Find(value)
	success := false

	if node2del != nil {
		prevNode := node2del.prev
		nextNode := node2del.next
		prevNode.next = node2del.next
		nextNode.prev = node2del.prev
		success = true
	}

	return success
}

func (m *DoubleLinkedList) PrintList() {
	for n := m.First(); n != nil; n = n.Next() {
		fmt.Printf("%d ", n.value)
	}
	fmt.Println("\n")
}

func main() {
	list := new(DoubleLinkedList)
	list.Push(5)
	list.Push(3)
	list.PrintList()
	list.Delete(3)
}
