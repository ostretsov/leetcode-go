package problems

// https://leetcode.com/problems/target-sum/description/
func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		if target == 0 {
			return 1
		}
		return 0
	}
	var waysToTarget int
	findTargetSumWaysLoop(nums, target, 0, &waysToTarget)
	return waysToTarget
}

func findTargetSumWaysLoop(nums []int, target, sum int, waysToTarget *int) {
	if len(nums) == 1 {
		if sum+nums[0] == target {
			*waysToTarget++
		}
		if sum-nums[0] == target {
			*waysToTarget++
		}
		return
	}
	findTargetSumWaysLoop(nums[1:], target, sum+nums[0], waysToTarget)
	findTargetSumWaysLoop(nums[1:], target, sum-nums[0], waysToTarget)
}
