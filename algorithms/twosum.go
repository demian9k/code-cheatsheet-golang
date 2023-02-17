package algorithms

func TwoSum(nums []int, target int) []int {
	indices := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := indices[complement]; ok {
			return []int{j, i}
		}
		indices[num] = i
	}
	return nil
}
