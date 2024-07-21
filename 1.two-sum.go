/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
    temp := make(map[int]int)
	var result []int
	for index, num := range nums {
		comp := target - num
		if val, found := temp[comp]; found {
			result = []int{val, index}
			return result
		}
		temp[num] = index
	}
	return result
}
// @lc code=end

