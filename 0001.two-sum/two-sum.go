package main

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for idx, num := range nums {
		if jdx, ok := index[target-num]; ok {
			return []int{jdx, idx}
		}
		index[num] = idx
	}
	return nil
}
