package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	type testCase struct {
		Name        string
		S           []int
		X           int
		ExpectedInt int
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualInt := Index(tc.S, tc.X)
			assert.Equal(t, tc.ExpectedInt, actualInt)
		})
	}

	validate(t, &testCase{
		Name: "first", S: []int{1, 1, 2, 3}, X: 3, ExpectedInt: 3,
	})
}

func TestCreatObject(t *testing.T) {
	CreateObjects()
}
