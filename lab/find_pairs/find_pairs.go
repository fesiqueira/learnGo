package main

import (
	"fmt"
	"sort"
)

type CallRecord struct {
	Type   string
	CallID int
}

type Call []CallRecord

func PairCalls(records []CallRecord) []Call {
	records = sortByCallID(records)
	calls := pairByCallID(records)

	return calls
}

func sortByCallID(records []CallRecord) []CallRecord {
	sort.Slice(records, func(i, j int) bool {
		return records[i].CallID < records[j].CallID
	})
	return records
}

func pairByCallID(records []CallRecord) []Call {
	calls := make([]Call, 0)

	for i := 0; i < len(records)-1; i++ {
		if records[i].CallID == records[i+1].CallID {
			calls = append(calls, Call(records[i:i+2]))
			i++
		}
	}
	return calls
}

func FindPairs(list []int) {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	fmt.Println(search(list, 10))
}

func search(list []int, x int) (int, bool) {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		if list[mid] > x {
			high = mid - 1
		} else if list[mid] < x {
			low = mid + 1
		} else {
			return mid, true
		}
	}
	return -1, false
}
