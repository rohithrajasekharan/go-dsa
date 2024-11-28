package main

type Node struct {
	value string
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) InsertFront(data string) {
	newNode := &Node{
		value: data,
		next:  ll.head,
	}
	ll.head = newNode
}

func (ll *LinkedList) InsertEnd(data string) {
	newNode := &Node{
		value: data,
		next:  nil,
	}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) InsertAfter(after, data string) bool {
	current := ll.head
	for current != nil && current.value != after {
		current = current.next
	}
	if current == nil {
		return false
	}
	newNode := &Node{
		value: data,
		next:  current.next,
	}
	current.next = newNode
	return true
}

func (ll *LinkedList) DeleteFront() bool {
	if ll.head == nil {
		return false
	}
	ll.head = ll.head.next
	return true
}

func (ll *LinkedList) Delete(data string) bool {
	if ll.head == nil {
		return false
	}
	if ll.head.value == data {
		ll.head = ll.head.next
		return true
	}
	current := ll.head
	for current.next != nil && current.next.value != data {
		current = current.next
	}
	if current.next == nil {
		return false
	} else {
		current.next = current.next.next
		return true
	}
}

func (ll *LinkedList) Print() {
	current := ll.head
	for current != nil {
		print(current.value, " -> ")
		current = current.next
	}
}

func main() {
	list := &LinkedList{}
	list.InsertFront("a")
	list.InsertEnd("z")
	list.InsertAfter("a", "b")

	//delete elements
	list.Delete("b")
	list.DeleteFront()
}
