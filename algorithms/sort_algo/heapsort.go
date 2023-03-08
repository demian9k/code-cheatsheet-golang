package sort_algo

import (
	"code-cheatsheet-golang/datastructure/heapq"
	"container/heap"
)

type sortableNum interface {
	int | int8 | int16 | int32 | int64
}

func heapsort[T sortableNum](iterable []T, maxHeap bool) []T {
	h := make(heapq.PriorityQueue[T], 0)
	result := make([]T, 0)

	// 모든 원소를 차례대로 힙에 삽입
	for _, value := range iterable {
		var v = value

		if maxHeap {
			v = -v
		}
		item := &heapq.Item[T]{Value: v, Priority: int(v)}
		heap.Push(&h, item)
	}

	// 힙에 삽입된 모든 원소를 꺼내기
	for len(h) > 0 {
		item := heap.Pop(&h).(*heapq.Item[T])
		v := item.Value
		if maxHeap {
			v = -v
		}
		result = append(result, v)
	}

	//if maxHeap {
	//	// 역순으로 뒤집어 최대 힙으로 변환
	//	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
	//		result[i], result[j] = result[j], result[i]
	//	}
	//}

	return result
}
