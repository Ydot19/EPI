package numberofStepsToReduceToZero_test

import (
	n "github.com/Ydot19/epi/leetcode/numberofStepsToReduceToZero"
	"github.com/stretchr/testify/suite"
	"testing"
)

type NumberOfStepsSuite struct {
	suite.Suite
	Impl n.NumberOfSteps
}

func (suite *NumberOfStepsSuite) SetupSuite() {
	suite.Impl = n.NewNumberOfStepsImpl()
}

func (suite *NumberOfStepsSuite) Cases(f func(int) int) {
	type testCase struct {
		description string
		num         int
		expected    int
	}

	for _, scenario := range []testCase{
		{
			description: "Base Case A",
			num:         14,
			expected:    6,
		},
		{
			description: "Base Case B",
			num:         8,
			expected:    4,
		},
		{
			description: "Base Case C",
			num:         123,
			expected:    12,
		},
		{
			description: "Base Case D",
			num:         9311,
			expected:    21,
		},
		{
			description: "Base Case E",
			num:         723119,
			expected:    29,
		},
	} {
		suite.T().Run(scenario.description, func(t *testing.T) {
			actual := f(scenario.num)
			suite.Equal(scenario.expected, actual)
		})
	}
}

func (suite *NumberOfStepsSuite) TestSolutionA() {
	suite.Cases(suite.Impl.SolutionA)
}

func (suite *NumberOfStepsSuite) TestOptimized() {
	suite.Cases(suite.Impl.Optimized)
}

func TestNumberOfSteps(t *testing.T) {
	suite.Run(t, new(NumberOfStepsSuite))
}
