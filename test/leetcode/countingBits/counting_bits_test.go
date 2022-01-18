package countingBits_test

import (
	"github.com/Ydot19/epi/leetcode/countingBits"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CountingBitsSuite struct {
	suite.Suite
	Impl countingBits.CountingBits
}

func (suite *CountingBitsSuite) SetupTest() {
	suite.Impl = countingBits.NewCountingBits()
}

func (suite *CountingBitsSuite) Cases(f func(n int) []int) {
	type testCase struct {
		description string
		n           int
		expected    []int
	}

	for _, scenario := range []testCase{
		{
			description: "Case A",
			n:           2,
			expected:    []int{0, 1, 1},
		},
		{
			description: "Case B",
			n:           5,
			expected:    []int{0, 1, 1, 2, 1, 2},
		},
		{
			description: "Case C",
			n:           98,
			expected: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4,
				4, 5, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5,
				4, 5, 5, 6, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4,
				4, 5, 4, 5, 5, 6, 2, 3, 3},
		},
		{
			description: "Case D",
			n:           211,
			expected: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4,
				4, 5, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5,
				4, 5, 5, 6, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4,
				4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6,
				4, 5, 5, 6, 5, 6, 6, 7, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4,
				4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5,
				4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4,
				4, 5},
		},
	} {
		suite.T().Run(scenario.description, func(t *testing.T) {
			actual := f(scenario.n)
			suite.Equal(scenario.expected, actual)
		})
	}
}

func (suite *CountingBitsSuite) TestSolutionA() {
	suite.Cases(suite.Impl.SolutionA)
}

func TestCountingBits(t *testing.T) {
	suite.Run(t, new(CountingBitsSuite))
}
