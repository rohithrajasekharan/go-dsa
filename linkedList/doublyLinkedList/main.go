package main

type Node struct {
  value interface{}
  next *Node
  prev *Node
}

type LinkedList struct {
  head *Node
}

func (ll *LinkedList) InsertFront(data string) {
  newNode := &Node{
    value:  data,
    prev: nil,
    next:  ll.head,
  }
  if ll.head == nil {
    ll.head = newNode
  }
}

func (ll *LinkedList) InsertAfter(after string, data string) bool {
  newNode := &Node{
    value: data,
    prev: nil,
    next: nil,
  }
  if ll.head.value == after {
    newNode.next = ll.head.next
    ll.head.next = newNode
    newNode.prev = ll.head
    return true 
  }
  current := ll.head
  for current!=nil && current.value!=after {
    current = current.next
  }
  if current != nil {
    newNode.next = current.next
    current.next = newNode
    current.prev = current
    return true
  }else{
    return false
  }
}

func (ll *LinkedList) InsertEnd(data string) {
  newNode := &Node{
    value: data,
    next: nil,
    prev: nil,
  }
  if ll.head == nil {
    ll.head = newNode
  }
  current := ll.head
  for current.next!=nil{
    current = current.next
  }
  newNode.prev = current
  current.next = newNode
}

func (ll *LinkedList) Delete(data string) bool {
  if ll.head == nil {
    return false
  }
  if ll.head.value == data {
    ll.head = ll.head.next
    ll.head.prev = nil
  }
  current := ll.head
  for current!=nil && current.value!=data {
    current = current.next
  }
  if current!= nil{
   current.next = current.next.next
   current.next.prev = current
   return true
 }else{
   return false
 }
} 

func (ll *LinkedList) DeleteAfter(data string) bool{
  if ll.head == nil{
    return false
  }
  if ll.head.value == data {
    ll.head.next = ll.head.next.next
    ll.head.next.prev = ll.head
    return true
  }
  current := ll.head
  for current!=nil && current.value!=data {
    current = current.next
  }
  if current!=nil{
    current.next = current.next.next
    current.next.prev = current
    return true
  }else{
    return false
  }
}

func (ll *LinkedList) DeleteFront() bool {
  if ll.head == nil {
    return false
  }
  ll.head = ll.head.next
  ll.head.prev = nil
  return true
}

func (ll *LinkedList) Print() {
	current := ll.head
	for current != nil {
		print(current.value, " -> ")
		current = current.next
	}
}
  
func main(){
	list := &LinkedList{}
	list.InsertFront("a")
	list.InsertEnd("z")
	list.InsertAfter("a", "b")

	//delete elements
	list.Delete("b")
	list.DeleteFront()

}
