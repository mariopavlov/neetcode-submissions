/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {

	var prev *ListNode
	cur := head

	for cur != nil {
		// Get the next pointer temporary so that we can move
		temp := cur.Next

		// Reverse the pointer
		cur.Next = prev

		// Move Pointers
		prev = cur
		cur = temp

	}

	return prev
}
