
type LinkedList struct {
	head *Node
}

type Node struct {
	val  int
	next *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
	cur := 0

	for node := ll.head; node != nil; node = node.next {
		if cur == index {
			return node.val
		}

		cur++
	}

	return -1
}

func (ll *LinkedList) InsertHead(val int) {
	newHead := &Node{val: val}

	newHead.next = ll.head

	ll.head = newHead
}

func (ll *LinkedList) InsertTail(val int) {
	newTail := &Node{val: val}

	// List Empty?
	if ll.head == nil {
		ll.head = newTail
		return
	}

	// Go to the end
	node := ll.head
	for node.next != nil {
		node = node.next
	}

	// Now at tail
	node.next = newTail
}

func (ll *LinkedList) Remove(index int) bool {
	// is empty?
	if ll.head == nil {
		return false
	}

	if index == 0 {
		ll.head = ll.head.next
		return true
	}

	cur := 0
	prev := ll.head

	for node := ll.head; node != nil; node = node.next {

		if cur == index {
			prev.next = node.next
			return true
		}

		prev = node
		cur++
	}

	return false
}

func (ll *LinkedList) GetValues() []int {
	if ll.head == nil {
		return []int{}
	}

	values := make([]int, 0)

	for node := ll.head; node != nil; node = node.next {
		values = append(values, node.val)
	}

	return values
}
