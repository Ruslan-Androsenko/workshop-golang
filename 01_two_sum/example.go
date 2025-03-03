package two_sum

func twoSum1(nums []int, target int) []int {
	newMap := make(map[int]int)

	for i, num := range nums {
		comp := target - num

		if index, found := newMap[comp]; found {
			return []int{index, i}
		}

		newMap[num] = i
	}

	return []int{}
}

func twoSum2(nums []int, target int) []int {
	hmap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		idx, ok := hmap[target-nums[i]]
		if ok {
			result := []int{idx, i}
			return result
		}
		hmap[nums[i]] = i
	}
	return []int{-1, -1}
}
