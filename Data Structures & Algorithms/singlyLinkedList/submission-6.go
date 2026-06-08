type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	val  int
	next *Node
}

// NewLinkedList returns an empty linked list.
func NewLinkedList() *LinkedList {
	node := &Node{}
	return &LinkedList{
		head: node,
		tail: node,
	}
}

// Get returns the value at the given index, or -1 if the index is out of range.
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

// InsertHead inserts a new node with val at the front of the list.
func (ll *LinkedList) InsertHead(val int) {
	node := &Node{
		val:  val,
		next: ll.head.next,
	}

	if node.next == nil {
		ll.tail = node
	}

	ll.head.next = node
}

// InsertTail inserts a new node with val at the end of the list.
func (ll *LinkedList) InsertTail(val int) {
	newTail := &Node{
		val: val,
	}

	ll.tail.next = newTail
	ll.tail = newTail
}

// Remove removes the node at the given index and returns true, or false if the index is invalid.
func (ll *LinkedList) Remove(index int) bool {
	i := 0

	for node := ll.head; node.next != nil; node = node.next {
		if i == index {
			if node.next == ll.tail {
				ll.tail = node
			}

			node.next = node.next.next
			return true
		}
		i++
	}

	return false
}

// GetValues returns all values in the list as a slice, in order.
func (ll *LinkedList) GetValues() []int {
	output := make([]int, 0)

	for node := ll.head.next; node != nil; node = node.next {
		output = append(output, node.val)
	}

	return output
}