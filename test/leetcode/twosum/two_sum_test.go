package twosum_test

import (
	"github.com/Ydot19/epi/leetcode/twosum"
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TwoSumSuite struct {
	suite.Suite
	Impl twosum.TwoSum
}

func (suite *TwoSumSuite) SetupTest() {
	suite.Impl = twosum.NewTwoSum()
}

func (suite *TwoSumSuite) Cases(a func(nums []int, target int) []int) {
	type testCases struct {
		description string
		nums        []int
		target      int
		expected    []int
	}
	for _, scenario := range []testCases{
		{
			description: "Base Case 1",
			nums:        []int{2, 7, 11, 15},
			target:      9,
			expected:    []int{0, 1},
		},
		{
			description: "Base Case 2",
			nums:        []int{3, 2, 4},
			target:      6,
			expected:    []int{1, 2},
		},
		{
			description: "Base Case 3",
			nums:        []int{3, 3},
			target:      6,
			expected:    []int{0, 1},
		},
	} {
		suite.T().Run(scenario.description, func(t *testing.T) {
			actual := a(scenario.nums, scenario.target)
			sort.Ints(actual)
			sort.Ints(scenario.expected)
			suite.Equal(scenario.expected, actual)
		})
	}

}

func (suite *TwoSumSuite) TestOptimized() {
	suite.Cases(suite.Impl.Optimized)
}

func TestTwoSum(t *testing.T) {
	suite.Run(t, new(TwoSumSuite))
}
