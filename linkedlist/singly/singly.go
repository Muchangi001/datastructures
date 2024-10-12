package singly

import (
	"fmt"
	"log"
)

type Node struct {
	Data string
	Next *Node
}

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
}

func (s *SinglyLinkedList) Init(h *Node, t *Node) {
	if s.Head != nil {
		log.Println("Error : failed to initialize singly linked list, pointer to head is not null")
		return
	}
	if s.Tail != nil {
		log.Println("Error : failed to initialize singly linked list, pointer to tail is not nil")
		return
	}
	s.Head = h
	s.Tail = t
}

func (s *SinglyLinkedList) GetSize() (count uint) {
	count = uint(0)
	for currNode := s.Head; currNode != nil; currNode = currNode.Next {
		count++
	}
	return
}

func (s *SinglyLinkedList) PrintList() {
	if s.Head == nil {
		log.Println("Error : failed to print singly linked list, pointer to head is nil. Make sure to initialize the list first")
		return
	}

	i := 0
	for currNode := s.Head ; currNode != nil; currNode = currNode.Next {
		fmt.Printf("Node [%v] -> %v\n", i, currNode.Data)
		i++
	}
}

func (s *SinglyLinkedList) Prepend(n *Node) {
	if n == nil {
		log.Println("Error : failed to prepend, pointer to node is nil")
		return
	}
	if s.Head == nil {
		s.Head = n
		return
	}
	if s.Head == s.Tail {
		n.Next = s.Head
		s.Head = n
		s.Tail = s.Head.Next
		return
	}
	n.Next = s.Head
	s.Head = n
}

func (s *SinglyLinkedList) Insert(n *Node, i uint) {
	if n == nil {
		log.Println("Error : failed to insert, pointer to node is nil")
		return
	}
	if i == 0 {
		s.Prepend(n)
	}
	if i > s.GetSize() {
		log.Println("Error : failed to insert node, index is out of range")
		return
	}
	currNode := s.Head
	for j := uint(0); j != i - 1; j++{
		currNode = currNode.Next
	}
	n.Next = currNode.Next
	currNode.Next = n
}

func (s *SinglyLinkedList) Append(n *Node) {
	if s.Head == nil {
		s.Head = n
		s.Tail = s.Head
		return
	}
	if s.Head == s.Tail {
		s.Tail = n
		s.Head.Next = s.Tail
		return
	}
	s.Tail.Next = n
	s.Tail = n
}

func (s *SinglyLinkedList) DeleteHead() {
	if s.Head == nil {
		log.Println("Error : failed to delete head, pointer to head is nil")
		return
	}
	s.Head = nil
}

func (s *SinglyLinkedList) DeleteNode(i uint) {
	if i == 0 {
		s.DeleteHead()
	}
	
	currNode := s.Head
	for j := uint(0); j != i - 1; j++ {
		currNode = currNode.Next
	}
	currNode.Next = currNode.Next.Next
}

func (s *SinglyLinkedList) DeleteTail() {
	if s.Tail == nil {
		log.Println("Error : failed to delete tail, pointer to tail is nil")
		return
	}
	currNode := s.Head;
	for ; currNode.Next != s.Tail; currNode = currNode.Next {
	}
	currNode.Next = nil
	s.Tail = currNode
}

func (s *SinglyLinkedList) GetNodePosition(n *Node) (i uint) {
	i = uint(0)
	for currNode := s.Head; currNode != n; currNode = currNode.Next {
		i++
	}
	return 
}