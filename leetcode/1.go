package leetcode

func twoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	tmp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := tmp[target-nums[i]]; ok {
			ret[0], ret[1] = tmp[target-nums[i]], i
			return ret
		} else {
			tmp[nums[i]] = i
		}
	}
	return ret
}
