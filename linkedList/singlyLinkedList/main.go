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
		next:  ll.head,
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
		next:  nil,
	}
	current.next = newNode
	return true
}

func (ll *LinkedList) Delete(data string) bool {
	if ll.head == nil {
		return false
	}
	if ll.head.value == data {
		ll.head = nil
		return true
	}
	ll.head = ll.head.next
	return true
}

func (ll *LinkedList) DeleteFront() bool {
	return false
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
