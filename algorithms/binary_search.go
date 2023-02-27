package algorithms

func BinarySearch(list []int, target int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2

		guess := list[mid]

		if guess == target {
			return mid
		}

		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}

	return -1
}
