package chxleetcode
/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{0, nil}
    temp := result
    flag := 0

    for l1 != nil || l2 != nil {
        sum := flag
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        if sum >= 10 {
            sum -= 10
            flag = 1
        } else {
            flag = 0
        }

        temp.Next = &ListNode{sum, nil}
        temp = temp.Next
    }

    if flag > 0 {
        temp.Next = &ListNode{flag, nil}
    }

    return result.Next
}
// @lc code=end

