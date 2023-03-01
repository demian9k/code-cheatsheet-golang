package algorithms

func BinarySearch(list []int, target int, low int, high int) int {
	//low := 0
	//high := len(list) - 1

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

func BinarySearch_recur(list []int, target int, start int, end int) int {
	if start > end {
		return -1
	}
	//fmt.Printf("low: %d high: %d\n", start, end)

	mid := (start + end) / 2
	//guess := list[mid]
	//fmt.Printf("low: %d mid: %d high: %d guess %d \n", start, mid, end, guess)
	if list[mid] == target {
		return mid
	}

	if list[mid] > target {
		end = mid - 1
	} else {
		start = mid + 1
	}

	return BinarySearch_recur(list, target, start, end)
}

//func factorial(n int) int {
//	if n == 1 {
//		return 1
//	}
//	return n * factorial(n-1)
//}
//
//func tailFact(n int) int {
//	return factT(n-1, n)
//}
//
//func factT(n, current int) int {
//	if n == 1 {
//		return current
//	}
//	return factT(n-1, n*current)
//}

func BinarySearchTailRecursion(list []int, target int) int {
	return BinarySearchTailRecursionT(list, target, 0, len(list)-1)
}

func BinarySearchTailRecursionT(list []int, target int, start int, end int) int {
	if start > end {
		return -1
	}
	//fmt.Printf("low: %d high: %d\n", start, end)

	mid := (start + end) / 2
	//guess := list[mid]
	//fmt.Printf("low: %d mid: %d high: %d guess %d \n", start, mid, end, guess)
	if list[mid] == target {
		return mid
	}

	if list[mid] > target {
		end = mid - 1
	} else {
		start = mid + 1
	}

	return BinarySearchTailRecursionT(list, target, start, end)
}
