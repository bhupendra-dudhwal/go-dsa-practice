package array

import (
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected float64
	}{
		{name: "Simple Case", nums: []int{1, 12, -5, -6, 50, 3}, k: 4, expected: 12.75},
		{name: "All Positive Numbers", nums: []int{1, 2, 3, 4, 5, 6}, k: 3, expected: 5.0},
		{name: "All Negative Numbers", nums: []int{-1, -2, -3, -4, -5}, k: 2, expected: -1.5},
		{name: "All Elements Same", nums: []int{5, 5, 5, 5, 5}, k: 3, expected: 5.0},
		{name: "Single Element", nums: []int{10}, k: 1, expected: 10.0},
		{name: "k equals Array Length", nums: []int{1, 2, 3, 4, 5}, k: 5, expected: 3.0},
		{name: "Mixed Positive and Negative Numbers", nums: []int{1, -2, 3, 4, -5, 6}, k: 3, expected: 1.667},
		{name: "Large Mixed Array", nums: []int{-1, 2, 3, 4, -5, 6, 7, -8}, k: 3, expected: 3.0},
		{name: "Large Array with Same Elements", nums: []int{1, 1, 1, 1, 1}, k: 100, expected: 1.0},
		{name: "Randomized Large Array", nums: []int{10, -100, 30, 50, -60, 70, 90, 120}, k: 4, expected: 55.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindMaxAverage(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("FindMaxAverage(%v,%d)=%f; expected %f", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

func BenchmarkFindMaxAverage(b *testing.B) {

}
