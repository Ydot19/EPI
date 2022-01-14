package countNumberOfConsistentStrings_test

import (
	"github.com/Ydot19/epi/leetcode/countNumberOfConsistentStrings"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CountNumberOfConsistentStringsSuite struct {
	suite.Suite
	Impl countNumberOfConsistentStrings.CountNumberOfConsistentLetters
}

func (suite *CountNumberOfConsistentStringsSuite) SetupTest() {
	suite.Impl = countNumberOfConsistentStrings.NewCountNumberOfConsistentLetters()
}

func (suite *CountNumberOfConsistentStringsSuite) Cases(f func(allowed string, words []string) int) {
	type testCases struct {
		description string
		allowed     string
		words       []string
		expected    int
	}

	for _, scenario := range []testCases{
		{
			description: "Base Case 1",
			allowed:     "ab",
			words:       []string{"ad", "bd", "aaab", "baa", "badab"},
			expected:    2,
		},
		{
			description: "Base Case 2",
			allowed:     "abc",
			words:       []string{"a", "b", "c", "ab", "ac", "bc", "abc", "abcd"},
			expected:    7,
		},
		{
			description: "Base Case 3",
			allowed:     "cad",
			words:       []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"},
			expected:    4,
		},
	} {
		suite.T().Run(scenario.description, func(t *testing.T) {
			actual := f(scenario.allowed, scenario.words)
			suite.Equal(scenario.expected, actual)
		})
	}
}

func (suite *CountNumberOfConsistentStringsSuite) TestSolutionA() {
	suite.Cases(suite.Impl.SolutionA)
}

func (suite *CountNumberOfConsistentStringsSuite) TestOptimized() {
	suite.Cases(suite.Impl.Optimized)
}

func TestCountNumberOfConsistentStrings(t *testing.T) {
	suite.Run(t, new(CountNumberOfConsistentStringsSuite))
}
