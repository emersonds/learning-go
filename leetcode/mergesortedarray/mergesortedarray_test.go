package main

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	type testCase struct {
		nums1 []int
		nums2 []int
		m     int
		n     int
		want  []int
	}

	var tests = map[string]testCase{
		"Example 1": {
			nums1: []int{1, 2, 3, 0, 0, 0},
			nums2: []int{2, 5, 6},
			m:     3,
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		"Example 2": {
			nums1: []int{1},
			nums2: []int{},
			m:     1,
			n:     0,
			want:  []int{1},
		},
		"Example 3": {
			nums1: []int{0},
			nums2: []int{1},
			m:     0,
			n:     1,
			want:  []int{1},
		},
		"Example 4": {
			nums1: []int{4,0,0,0,0,0},
			nums2: []int{1,2,3,4,5,6},
			m: 1,
			n: 5,
			want: []int{1,2,3,4,5,6},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := merge(tc.nums1, tc.m, tc.nums2, tc.n)

			if s := slices.Equal(got, tc.want); !s {
				t.Errorf("Expected: %d, Got: %d", tc.want, got)
			}
		})
	}
}
