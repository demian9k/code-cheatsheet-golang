package perf

func Fibo(x int) int {
	if x == 1 || x == 2 {
		return 1
	}
	return Fibo(x-1) + Fibo(x-2)
}

/*
memoization 을 사용한 fibonacci 함수
*/
func ProduceFiboMemo() func(int) int {
	d := make([]int, 100)

	var fibo_memo func(int) int
	fibo_memo = func(x int) int {
		if x == 1 || x == 2 {
			return 1
		}

		// 전에 계산한 값은 계산없이 반환
		if d[x] != 0 {
			return d[x]
		}

		// 계산하지 않은 문제이면 점화식에 따라 계산 후 반환
		d[x] = fibo_memo(x-1) + fibo_memo(x-2)

		return d[x]
	}

	return fibo_memo
}

/*
tail recursion을 사용한 fibonacci 함수
*/
func fiboHelper(x, acc1, acc2 int) int {
	if x == 1 {
		return acc1
	} else if x == 2 {
		return acc2
	} else {
		return fiboHelper(x-1, acc2, acc1+acc2)
	}
}

func FiboTr(x int) int {
	return fiboHelper(x, 1, 1)
}

/*
tail recursion + memoization을 사용한 fibonacci 함수
*/
func produceFiboTrMemo() func(int, int, int) int {
	d := make([]int, 100)

	var fibo_tr_memo_helper func(int, int, int) int

	fibo_tr_memo_helper = func(x int, acc1 int, acc2 int) int {
		if x == 1 {
			return acc1
		} else if x == 2 {
			return acc2
		} else {
			// 전에 계산한 값은 계산없이 반환
			if d[x] != 0 {
				return d[x]
			}

			d[x] = fibo_tr_memo_helper(x-1, acc2, acc1+acc2)

			return d[x]
		}
	}

	return fibo_tr_memo_helper
}

func FiboTrMemo(x int) int {
	fibo_tr_memo_helper := produceFiboTrMemo()
	return fibo_tr_memo_helper(x, 1, 1)
}
