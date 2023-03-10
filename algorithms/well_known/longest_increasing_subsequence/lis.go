package algorithms

/*
GetLongestIncreasingSubsequence

	DP를 사용해 LIS 의 길이를 구한다.
*/
func GetLongestIncreasingSubsequence(arr []int) int {
	n := len(arr)
	lenList := make([]int, n)
	for i := 0; i < n; i++ {
		lenList[i] = 1
	}

	for k := 0; k < n; k++ {
		for i := 0; i < k; i++ {
			if arr[i] < arr[k] {
				lenList[k] = max(lenList[k], lenList[i]+1)
			}
		}
	}

	maxLen := 1
	for _, l := range lenList {
		if l > maxLen {
			maxLen = l
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
LIS

	LIS 구조체는 배열의 최장 증가 부분 수열(LIS)을 구하는 데 사용합니다.
	각 원소를 저장하고, LIS 길이와 각 원소의 인덱스, LIS를 저장하는 배열 등을 관리합니다.
*/
type LIS struct {
	arr   []int
	n     int
	L     []int
	index []int
	ans   []int
}

// 입력으로 받은 배열 arr을 저장하고, n, L, index, ans 배열을 초기화합니다.
func NewLIS(arr []int) *LIS {
	n := len(arr)
	L := make([]int, n)
	index := make([]int, n)
	ans := make([]int, n)
	return &LIS{arr, n, L, index, ans}
}

/*
GetLISLength

 LIS의 길이를 반환합니다.
 L, index 배열을 갱신하면서 이진 탐색(lower bound)을 이용하여 LIS의 길이를 계산합니다.
*/
func (l *LIS) GetLISLength() int {
	idx := 0
	for i := 0; i < l.n; i++ {
		temp := l.arr[i]
		if idx == 0 {
			l.L[idx] = temp
			l.index[i] = 0 // 첫 번째 위치는 0
			idx++
		} else {
			if l.L[idx-1] < temp {
				l.index[i] = idx
				l.L[idx] = temp
				idx++
			} else {
				l.index[i] = l.lowerBound(temp, idx)
				l.L[l.lowerBound(temp, idx)] = temp
			}
		}
	}
	return idx
}

/*
lowerBound

이진 탐색(lower bound)을 이용하여 target 값의 위치를 반환합니다.
L 배열에서 target 값보다 처음으로 크거나 같은 값이 나오는 위치를 찾습니다.
*/
func (l *LIS) lowerBound(target, end int) int {
	start := 0
	for start < end {
		mid := (start + end) / 2
		if l.L[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return end
}

/*
GetLIS

LIS를 반환합니다.
index 배열을 이용하여 LIS에 포함되는 원소의 인덱스를 찾고, 역순으로 ans 배열에 저장합니다.
ans 배열의 원소들을 역순으로 출력하면 LIS가 됩니다.
*/
func (l *LIS) GetLIS() []int {
	idx := l.GetLISLength()
	t := 0
	for i := l.n - 1; i >= 0; i-- {
		if idx == l.index[i]+1 {
			l.ans[t] = l.arr[i]
			idx--
			t++
		}
	}
	return l.ans[:t]
}
