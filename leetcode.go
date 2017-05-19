package leetcode

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j, w := range nums {
			if i == j {
				continue
			}
			if v+w == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
