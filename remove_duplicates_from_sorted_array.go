package leetcode

// Given a sorted array, remove the duplicates in-place such that each element
// appear only once and return the new length.
// Do not allocate extra space for another array, you must do this by modifying
// the input array in-place with O(1) extra memory.
// 6 minutes implementation time
func RemoveDuplicatesFromSortedArray(nums []int) int {
	dupes := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			dupes++
		} else {
			nums[i-dupes] = nums[i]
		}
	}
	return len(nums) - dupes
}
