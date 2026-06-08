type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	val  int
	next *Node
}

func NewLinkedList() *LinkedList {
	head := &Node{val: -1}
	return &LinkedList{
		head: head,
		tail: head,
	}
}

func (ll *LinkedList) Get(index int) int {
	cur := 0

	for node := ll.head; node != nil; node = node.next {

		if node.val == -1 {
			continue
		}

		if cur == index {
			return node.val
		}

		cur++
	}

	return -1
}

func (ll *LinkedList) InsertHead(val int) {
	newHead := &Node{val: val}

	if ll.head.next == nil {
		ll.tail = newHead
	}

	newHead.next = ll.head.next

	ll.head.next = newHead
}

func (ll *LinkedList) InsertTail(val int) {
	tail := &Node{val: val}

	ll.tail.next = tail
	ll.tail = tail
}

func (ll *LinkedList) Remove(index int) bool {
	cur := 0
	prev := ll.head

	for node := ll.head; node != nil; node = node.next {
		if node.val == -1 {
			continue
		}

		if cur == index {
			prev.next = node.next

			if prev.next == nil {
				ll.tail = prev
			}

			return true
		}

		prev = node
		cur++
	}

	return false
}

func (ll *LinkedList) GetValues() []int {
	values := make([]int, 0)

	for node := ll.head; node != nil; node = node.next {
		if node.val == -1 {
			continue
		}
		values = append(values, node.val)
	}

	return values
}
