package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCases struct {
	Name     string
	Actual   interface{}
	Expected interface{}
	Data
}

type Data struct {
	Number1 int64
	Number2 int64
}

func TestAddition(t *testing.T) {
	testCases := []TestCases{
		{
			Name: "Success - Normal Addition 1",
			Data: Data{
				Number1: 10,
				Number2: 2,
			},
			Expected: int64(12),
		},
		{
			Name: "Success - Normal Addition 2",
			Data: Data{
				Number1: 1,
				Number2: 2,
			},
			Expected: int64(3),
		},
		{
			Name: "Success - Normal Addition 3",
			Data: Data{
				Number1: -10,
				Number2: 2,
			},
			Expected: int64(-8),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Actual = Addition(tc.Number1, tc.Number2)

			assert.Equal(t, tc.Expected, tc.Actual)
		})
	}
}
