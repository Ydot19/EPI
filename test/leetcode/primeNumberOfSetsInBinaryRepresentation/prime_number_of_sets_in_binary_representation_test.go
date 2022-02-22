package primeNumberOfSetsInBinaryRepresentation_test

import (
	p "github.com/Ydot19/epi/leetcode/primeNumberOfSetsInBinaryRepresentation"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestPrimeNumberOfSetsInBinaryRepresentation(t *testing.T) {
	suite.Run(t, new(PrimeNumberOfSetsInBinaryRepresentationSuite))
}

type PrimeNumberOfSetsInBinaryRepresentationSuite struct {
	suite.Suite
	Impl p.PrimeNumberOfSetsInBinaryRepresentation
}

func (suite *PrimeNumberOfSetsInBinaryRepresentationSuite) SetupTest() {
	suite.Impl = p.NewPrimeNumberOfSetsInBinaryRepresentation()
}

func (suite *PrimeNumberOfSetsInBinaryRepresentationSuite) Cases(f func(left, right int) int) {
	type testCase struct {
		description string
		left        int
		right       int
		expected    int
	}

	for _, scenario := range []testCase{
		{
			description: "Case 1",
			left:        6,
			right:       10,
			expected:    4,
		},
		{
			description: "Case 2",
			left:        10,
			right:       15,
			expected:    5,
		},
		{
			description: "Case 3",
			left:        14,
			right:       97,
			expected:    54,
		},
		{
			description: "Case 4",
			left:        131,
			right:       926,
			expected:    416,
		},
	} {
		suite.T().Run(scenario.description, func(t *testing.T) {
			actual := f(scenario.left, scenario.right)
			suite.Equal(scenario.expected, actual)
		})
	}
}

func (suite *PrimeNumberOfSetsInBinaryRepresentationSuite) TestOptimized() {
	suite.Cases(suite.Impl.Optimized)
}
