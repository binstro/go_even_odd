package goevenodd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenSuccess(t *testing.T) {
	result := CheckEvenOdd(2)
	assert.Equal(t, "genap", result, "Test result must be 'genap'")
}

func TestOddFailed(t *testing.T) {
	result := CheckEvenOdd(1)
	assert.Equal(t, "ganjil", result, "Test result must be 'ganjil'")
}

func TestEvenOddWithTable(t *testing.T) {
	tests := []struct {
		name     string
		request  int
		expected string
	}{
		{
			name:     "Check even number",
			request:  2,
			expected: "genap",
		}, {
			name:     "Check odd number",
			request:  1,
			expected: "ganjil",
		}, {
			name:     "Check zero",
			request:  0,
			expected: "genap",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CheckEvenOdd(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
