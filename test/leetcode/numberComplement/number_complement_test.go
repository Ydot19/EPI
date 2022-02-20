package numberComplement_test

import (
	"github.com/Ydot19/epi/leetcode/numberComplement"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestNumberComplement(t *testing.T) {
	suite.Run(t, new(NumberComplementTestSuite))
}

type NumberComplementTestSuite struct {
	suite.Suite
	Impl numberComplement.NumberComplement
}

func (suite *NumberComplementTestSuite) SetupTest() {
	suite.Impl = numberComplement.NewNumberComplement()
}

func (suite *NumberComplementTestSuite) Cases(f func(num int) int) {
	type testCase struct {
		description string
		num         int
		expected    int
	}

	for _, scenario := range []testCase{
		{
			description: "Case 1",
			num:         5,
			expected:    2,
		},
		{
			description: "Case 2",
			num:         2147483646,
			expected:    1,
		},
		{
			description: "Case 3",
			num:         319,
			expected:    192,
		},
	} {
		suite.T().Run(scenario.description, func(t *testing.T) {
			actual := f(scenario.num)
			suite.Equal(scenario.expected, actual)
		})
	}
}

func (suite *NumberComplementTestSuite) TestSolutionA() {
	suite.Cases(suite.Impl.SolutionA)
}

func (suite *NumberComplementTestSuite) TestOptimized() {
	suite.Cases(suite.Impl.Optimized)
}
