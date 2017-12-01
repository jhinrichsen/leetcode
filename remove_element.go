package leetcode

func removeElement(nums []int, val int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			n++
		} else {
			nums[i-n] = nums[i]
		}
	}
	return len(nums) - n
}
