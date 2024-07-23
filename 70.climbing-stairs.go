/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 1 { 
		return 1
	}
	num1 := 1
	num2 := 2
	for range n-2 {
		temp := num2
		num2 += num1
		num1 = temp
	}
	return num2
	// 1: 1 = 1 
	// 2: 2 = 1 1 | 2 = 2
	// 3: 3 = 1 1 1 | 2 1 | 1 2 = 3
	// 4: 5 = 1 1 1 1 | 2 1 1 | 1 2 1 | 1 1 2 | 2 2 = 5
	// 5: 8 = 1 1 1 1 1 | 2 1 1 1 | 1 2 1 1 | 1 1 2 1 | 1 1 1 2 | 2 2 1 |  2 1 2 | 1 2 2
}
// @lc code=end

