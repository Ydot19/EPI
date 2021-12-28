package decodedXORedArray_test

import (
	"github.com/Ydot19/epi/leetcode/decodedXORedArray"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DecodedXORedArraySuite struct {
	suite.Suite
	Impl decodedXORedArray.DecodedXORedArray
}

func (suite *DecodedXORedArraySuite) SetupTest() {
	suite.Impl = decodedXORedArray.NewDecodedXORedArrayImpl()
}

func (suite *DecodedXORedArraySuite) Cases(f func([]int, int) []int) {
	type testCases struct {
		description string
		encoded     []int
		first       int
		expected    []int
	}

	for _, scenario := range []testCases{
		{
			description: "Base Case 1",
			encoded:     []int{1, 2, 3},
			first:       1,
			expected:    []int{1, 0, 2, 1},
		},
		{
			description: "Base Case 2",
			encoded:     []int{6, 2, 7, 3},
			first:       4,
			expected:    []int{4, 2, 0, 7, 4},
		},
	} {
		suite.T().Run(scenario.description, func(t *testing.T) {
			actual := f(scenario.encoded, scenario.first)
			suite.Equal(scenario.expected, actual)
		})
	}
}

func (suite *DecodedXORedArraySuite) TestOptimized() {
	suite.Cases(suite.Impl.Optimized)
}

func TestDecodedXORedArray(t *testing.T) {
	suite.Run(t, new(DecodedXORedArraySuite))
}
