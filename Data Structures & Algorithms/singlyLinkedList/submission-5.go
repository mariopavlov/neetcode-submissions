type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	val  int
	next *Node
}

func NewLinkedList() *LinkedList {
	newHead := &Node{val: 1}

	return &LinkedList{
		head: newHead,
		tail: newHead,
	}
}

func (ll *LinkedList) Get(index int) int {
	i := 0
	for node := ll.head.next; node != nil; node = node.next {
		if i == index {
			return node.val
		}
		i++
	}

	return -1
}

func (ll *LinkedList) InsertHead(val int) {
	newNode := &Node{val: val}

	newNode.next = ll.head.next
	ll.head.next = newNode

	if newNode.next == nil {
		ll.tail = newNode
	}
}

func (ll *LinkedList) InsertTail(val int) {
	newNode := &Node{val: val}

	ll.tail.next = newNode
	ll.tail = newNode
}

func (ll *LinkedList) Remove(index int) bool {
	if index < 0 {
		return false
	}

	i := 0
	node := ll.head

	for i < index && node != nil {
		i++
		node = node.next
	}

	if node != nil && node.next != nil {

		if node.next == ll.tail {
			ll.tail = node
		}
		node.next = node.next.next

		return true
	}

	return false
}

func (ll *LinkedList) GetValues() []int {
	out := make([]int, 0)

	for node := ll.head.next; node != nil; node = node.next {
		out = append(out, node.val)
	}

	return out
}
