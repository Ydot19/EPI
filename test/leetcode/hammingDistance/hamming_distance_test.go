package hammingDistance_test

import (
	"github.com/Ydot19/epi/leetcode/hammingDistance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HammingDistanceSuite struct {
	suite.Suite
	Impl hammingDistance.HammingDistance
}

func (suite *HammingDistanceSuite) SetupSuite() {
	suite.Impl = hammingDistance.NewHammingDistance()
}

func (suite *HammingDistanceSuite) Cases(f func(int, int) int) {
	type testCase struct {
		description string
		x           int
		y           int
		expected    int
	}

	for _, scenario := range []testCase{
		{
			description: "Base Case 1",
			x:           1,
			y:           4,
			expected:    2,
		},
		{
			description: "Base Case 2",
			x:           10,
			y:           99,
			expected:    4,
		},
		{
			description: "Base Case 3",
			x:           1093,
			y:           93443112,
			expected:    13,
		},
		{
			description: "Base Case 4",
			x:           749199,
			y:           213,
			expected:    12,
		},
		{
			description: "Base Case 5",
			x:           5639,
			y:           1436241892,
			expected:    18,
		},
	} {
		suite.T().Run(scenario.description, func(t *testing.T) {
			actual := f(scenario.x, scenario.y)
			assert.Equal(t, scenario.expected, actual)
		})
	}
}

func (suite *HammingDistanceSuite) TestOptimized() {
	suite.Cases(suite.Impl.Optimized)
}

func TestHammingDistance(t *testing.T) {
	suite.Run(t, new(HammingDistanceSuite))
}
