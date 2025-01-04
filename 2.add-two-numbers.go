// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		head    *ListNode
		current *ListNode
		carry   int
	)

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		newNode := &ListNode{Val: sum % 10}

		if head == nil {
			head = newNode
			current = newNode
		} else {
			current.Next = newNode
			current = newNode
		}
	}

	return head
}

// @leet end
