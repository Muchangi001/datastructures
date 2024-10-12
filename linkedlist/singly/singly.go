package singly

import (
	"fmt"
	"log"
)

type Node struct {
	data string
	next *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
}

func (s *SinglyLinkedList) Init(h *Node, t *Node) {
	if s.head != nil {
		log.Println("Error : failed to initialize singly linked list, pointer to head is not null")
		return
	}
	if s.tail != nil {
		log.Println("Error : failed to initialize singly linked list, pointer to tail is not nil")
		return
	}
	s.head = h
	s.tail = t
}

func (s *SinglyLinkedList) GetSize() (count uint) {
	count = uint(0)
	for currNode := s.head; currNode != nil; currNode = currNode.next {
		count++
	}
	return
}

func (s *SinglyLinkedList) PrintList() {
	if s.head == nil {
		log.Println("Error : failed to print singly linked list, pointer to head is nil. Make sure to initialize the list first")
		return
	}

	i := 0
	for currNode := s.head ; currNode != nil; currNode = currNode.next {
		fmt.Printf("Node [%v] -> %v\n", i, currNode.data)
		i++
	}
}

func (s *SinglyLinkedList) Prepend(n *Node) {
	if n == nil {
		log.Println("Error : failed to prepend, pointer to node is nil")
		return
	}
	if s.head == nil {
		s.head = n
		return
	}
	if s.head == s.tail {
		n.next = s.head
		s.head = n
		s.tail = s.head.next
		return
	}
	n.next = s.head
	s.head = n
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
	currNode := s.head
	for j := uint(0); j != i - 1; j++{
		currNode = currNode.next
	}
	n.next = currNode.next
	currNode.next = n
}

func (s *SinglyLinkedList) Append(n *Node) {
	if s.head == nil {
		s.head = n
		s.tail = s.head
		return
	}
	if s.head == s.tail {
		s.tail = n
		s.head.next = s.tail
		return
	}
	s.tail.next = n
	s.tail = n
}

func (s *SinglyLinkedList) DeleteHead() {
	if s.head == nil {
		log.Println("Error : failed to delete head, pointer to head is nil")
		return
	}
	s.head = nil
}

func (s *SinglyLinkedList) DeleteNode(i uint) {
	if i == 0 {
		s.DeleteHead()
	}
	
	currNode := s.head
	for j := uint(0); j != i - 1; j++ {
		currNode = currNode.next
	}
	currNode.next = currNode.next.next
}

func (s *SinglyLinkedList) DeleteTail() {
	if s.tail == nil {
		log.Println("Error : failed to delete tail, pointer to tail is nil")
		return
	}
	currNode := s.head;
	for ; currNode.next != s.tail; currNode = currNode.next {
	}
	currNode.next = nil
	s.tail = currNode
}

func (s *SinglyLinkedList) GetNodePosition(n *Node) (i uint) {
	i = uint(0)
	for currNode := s.head; currNode != n; currNode = currNode.next {
		i++
	}
	return 
}