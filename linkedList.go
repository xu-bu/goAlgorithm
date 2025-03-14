package main

import "fmt"

type node struct {
	val  int
	next *node
}

type linkedList struct {
	head *node
}

func (this *linkedList) traverse() {
	p := this.head
	for {
		fmt.Println(p.val)
		if p.next != nil {
			p = p.next
		} else {
			break
		}
	}
}

func (this *linkedList) length() int {
	count := 0
	p := this.head
	for {
		count++
		if p.next != nil {
			p = p.next
		} else {
			break
		}
	}
	return count
}

func (this *linkedList) join(list *linkedList) {
	p := this.head
	nextHead := list.head
	for {
		if p.next != nil {
			p = p.next
		} else {
			break
		}
	}
	p.next = nextHead
}

func (this *linkedList) listCut(position int) (*linkedList, *linkedList) {
	n := this.length()
	if n <= 1 || position >= n {
		return this, nil
	} else {
		newHead := this.head
		var pre *node
		for position > 0 && newHead != nil {
			pre = newHead
			newHead = newHead.next
			position--
		}
		pre.next = nil
		newList := &linkedList{newHead}
		return this, newList
	}
}

func (this *linkedList) get(position int) *node {
	n := this.length()
	if position >= n {
		fmt.Println("out of range")
		return &node{-1, nil}
	} else {
		p := this.head
		for position > 0 && p != nil {
			p = p.next
			position--
		}
		return p
	}
}

func (this *linkedList) tailAppend(val int) {
	//不要新插入一个节点，因为如果新节点后面还有节点，会把后面的节点也插进来
	newNode := &node{val, nil}
	if this.head == nil {
		this.head = newNode
	} else {
		p := this.head
		for p.next != nil {
			p = p.next
		}
		p.next = newNode
	}
}

