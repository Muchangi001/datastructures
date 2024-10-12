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

func (s *SinglyLinkedList) init(h *Node, t *Node) {
	if s.head != nil {
		log.Println("Error : failed to initialize singly linked list, pointer to head is not null")
		return
	}
	if s.tail != nil {
		log.Println("Error : failed to initialize singly linked list, pointer to tail is not null")
		return
	}
	s.head = h
	s.tail = t
}

func (s *SinglyLinkedList) printList() {
	if s.head == nil {
		log.Println("Error : failed print singly linked list, pointer to head is null. Make sure to initialize the list first")
		return
	}

	i := 0
	for currentNode := s.head ; currentNode.next != nil; currentNode = currentNode.next {
		fmt.Printf("Node [%v] -> %v", i, currentNode.data)
		i++
	}
}

func (s *SinglyLinkedList) prepend(n *Node) {
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