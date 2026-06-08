/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {

	var cur *ListNode
	next := head

	for next != nil {
		temp := next.Next

		next.Next = cur

		cur = next
		next = temp
	}

	return cur
}
