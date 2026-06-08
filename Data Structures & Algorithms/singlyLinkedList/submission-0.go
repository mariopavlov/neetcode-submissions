
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
	node := Node{val: val}

	head := ll.head
	ll.head = &node
	ll.head.next = head
}

func (ll *LinkedList) InsertTail(val int) {
	last := ll.getLastNode()

	if last == nil {
		ll.head = &Node{val: val}
		return
	}

	last.next = &Node{val: val}
}

func (ll *LinkedList) Remove(index int) bool {
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

		cur++
		prev = node
	}

	return false
}

func (ll *LinkedList) GetValues() []int {
	values := make([]int, 0)

	for node := ll.head; node != nil; node = node.next {
		values = append(values, node.val)
	}

	return values
}

func (ll *LinkedList) getLastNode() *Node {
	if ll.head == nil {
		return nil
	}

	node := ll.head

	for node.next != nil {
		node = node.next
	}

	return node
}
