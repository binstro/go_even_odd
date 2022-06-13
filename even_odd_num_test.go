package goevenodd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenSuccess(t *testing.T) {
	result := CheckEvenOdd(1, 2, 3, 4, 5)
	assert.Equal(t, "ganjil, genap, ganjil, genap, ganjil", result)
}

func TestEvenOddWithTable(t *testing.T) {
	tests := []struct {
		name     string
		request  []int
		expected string
	}{
		{
			name:     "Check Genap",
			request:  []int{0, 2, 4},
			expected: "genap, genap, genap",
		}, {
			name:     "Check Ganjil",
			request:  []int{1, 3, 5},
			expected: "ganjil, ganjil, ganjil",
		}, {
			name:     "Check Genap Ganjil",
			request:  []int{1, 2, 3, 4, 5},
			expected: "ganjil, genap, ganjil, genap, ganjil",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CheckEvenOdd(test.request...)
			assert.Equal(t, test.expected, result)
		})
	}
}
