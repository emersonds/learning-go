package main

import "testing"

func TestRemoveElement(t *testing.T) {
	type testCase struct {
		nums []int
		val  int
		want int
	}

	tests := map[string]testCase{
		"Test1": {
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: 2,
		},
		"Test2": {
			nums: []int{0,1,2,2,3,0,4,2},
			val: 2,
			want: 5,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := removeElement(tc.nums, tc.val)

			if got != tc.want {
				t.Errorf("Expected: %d, Got: %d", tc.want, got)
			}
		})
	}
}
