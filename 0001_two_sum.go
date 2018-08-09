package main

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		wantNum := target - nums[i]
		if index, ok := m[wantNum]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return nil
}
