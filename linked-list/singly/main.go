package main

import "fmt"

type node struct {
	val  int
	next *node
}

type singlyList struct {
	len  int
	head *node
	// tail *node
}

func initList() *singlyList {
	return &singlyList{}
}

func (s *singlyList) InsertAtHead(value int) {
	node := &node{
		val: value,
	}

	if s.head == nil {
		s.head = node
	} else {
		node.next = s.head
		s.head = node
	}

	s.len++
	return
}

func (s *singlyList) InsertAtTail(value int) {
	node := &node{
		val: value,
	}

	if s.head == nil {
		s.head = node
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}

		current.next = node
	}

	s.len++
	return
}

func (s *singlyList) GetNodeAtIndex(index int) *node {
	if index < 0 || index >= s.len {
		return nil
	}

	if index == 0 {
		return s.head
	}

	if index > 0 && index < s.len {
		current := s.head

		for current.next == nil {
			current = current.next
		}

		return current
	}

	return nil
}

func (s *singlyList) InsertAtIndex(index int, value int) {
	if index < 0 {
		return
	}

	if index == 0 {
		s.InsertAtHead(value)
		return
	}

	if index >= s.len {
		s.InsertAtTail(value)
		return
	}

	prev := s.GetNodeAtIndex(index - 1)
	node := &node{
		val:  value,
		next: prev.next,
	}

	prev.next = node

	s.len++
}

func main() {
	singleList := initList()
	singleList.InsertAtIndex(0, 1)
	singleList.InsertAtIndex(1, 2)
	fmt.Println(singleList)
}
