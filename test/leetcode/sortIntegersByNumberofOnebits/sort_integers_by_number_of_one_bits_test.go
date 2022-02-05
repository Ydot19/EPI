package sortIntegersByNumberofOnebits_test

import (
	sortIntegersByNumberOfOneDigits "github.com/Ydot19/epi/leetcode/sortIntegersByNumberOfOneBits"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SortIntegersByNumberOfOneBitsSuite struct {
	suite.Suite
	Impl sortIntegersByNumberOfOneDigits.SortIntegersByNumberOfOneBits
}

func (suite *SortIntegersByNumberOfOneBitsSuite) SetupTest() {
	suite.Impl = sortIntegersByNumberOfOneDigits.NewSortIntegersByNumberOfOneBits()
}

func (suite *SortIntegersByNumberOfOneBitsSuite) Cases(a func(arr []int) []int) {
	type testCase struct {
		description string
		arr         []int
		expected    []int
	}

	for _, scenario := range []testCase{
		{
			description: "Base Case 1",
			arr:         []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			expected:    []int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
		{
			description: "Base Case 2",
			arr:         []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
			expected:    []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
		},
	} {
		suite.T().Run(scenario.description, func(t *testing.T) {
			actual := a(scenario.arr)
			suite.Equal(scenario.expected, actual)
		})
	}
}

func (suite *SortIntegersByNumberOfOneBitsSuite) TestSolutionA() {
	suite.Cases(suite.Impl.SolutionA)
}

func TestSortIntegersByNumberOfOneBitsSuite(t *testing.T) {
	suite.Run(t, new(SortIntegersByNumberOfOneBitsSuite))
}
