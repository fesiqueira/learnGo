package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in  []CallRecord
	out []Call
}{
	{
		in:  []CallRecord{CallRecord{"end", 12}, CallRecord{"start", 10}, CallRecord{"start", 12}, CallRecord{"end", 10}, CallRecord{"start", 5}, CallRecord{"end", 5}},
		out: []Call{Call{CallRecord{"start", 10}, CallRecord{"end", 10}}, Call{CallRecord{"end", 12}, CallRecord{"start", 12}}, Call{CallRecord{"start", 5}, CallRecord{"end", 5}}},
	},
}

func TestSortByCallID(t *testing.T) {
	var cases = []struct {
		in  []CallRecord
		out []CallRecord
	}{
		{
			in:  []CallRecord{CallRecord{"end", 12}, CallRecord{"start", 10}},
			out: []CallRecord{CallRecord{"start", 10}, CallRecord{"end", 12}},
		},
		{
			in:  []CallRecord{CallRecord{"end", 12}, CallRecord{"start", 10}, CallRecord{"start", 12}, CallRecord{"end", 10}, CallRecord{"end", 5}},
			out: []CallRecord{CallRecord{"end", 5}, CallRecord{"start", 10}, CallRecord{"end", 10}, CallRecord{"end", 12}, CallRecord{"start", 12}},
		},
		{
			in:  []CallRecord{CallRecord{"end", 12}, CallRecord{"start", 10}, CallRecord{"start", 12}, CallRecord{"end", 10}, CallRecord{"start", 5}, CallRecord{"end", 5}},
			out: []CallRecord{CallRecord{"start", 5}, CallRecord{"end", 5}, CallRecord{"start", 10}, CallRecord{"end", 10}, CallRecord{"end", 12}, CallRecord{"start", 12}},
		},
	}

	for _, test := range cases {
		records := sortByCallID(test.in)
		assert.Equal(t, test.out, records)
	}
}
func TestPairByCallID(t *testing.T) {
	var cases = []struct {
		in  []CallRecord
		out []Call
	}{
		{
			in:  []CallRecord{CallRecord{"end", 12}, CallRecord{"start", 10}},
			out: []Call{},
		},
		{
			in:  []CallRecord{CallRecord{"end", 12}, CallRecord{"start", 10}, CallRecord{"start", 12}, CallRecord{"end", 10}, CallRecord{"end", 5}},
			out: []Call{Call{CallRecord{"start", 10}, CallRecord{"end", 10}}, Call{CallRecord{"end", 12}, CallRecord{"start", 12}}},
		},
	}

	for _, test := range cases {
		records := sortByCallID(test.in)
		calls := pairByCallID(records)
		assert.Equal(t, test.out, calls)
	}
}

func TestPairCalls(t *testing.T) {
	for _, test := range tests {
		assert.ElementsMatch(t, test.out, PairCalls(test.in))
	}
}

func BenchmarkPairCalls(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, test := range tests {
			_ = PairCalls(test.in)
		}
	}
}

// func BenchmarkFindPairs(b *testing.B) {
// 	// numbers := []int{1, 2, 3, 4, 5, 6, 7}
// 	for n := 0; n < b.N; n++ {
// 		FindPairs(numbers)
// 	}
// }
