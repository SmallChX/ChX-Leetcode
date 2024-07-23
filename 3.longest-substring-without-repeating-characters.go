/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[rune]int)
	maxLength := 0
	start := 0

	for i, c := range s {
		if index, found := charIndexMap[c]; found && index >= start {
			start = index + 1
		}
		charIndexMap[c] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	return maxLength
}
// @lc code=end

