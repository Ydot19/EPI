package countNumberOfConsistentStrings

import "github.com/Ydot19/epi/leetcode/utils"

//You are given a string allowed consisting of distinct characters and an array of strings words. A string is consistent if all characters in the string appear in the string allowed.
//
//Return the number of consistent strings in the array words.
//
//
//
//Example 1:
//
//Input: allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
//Output: 2
//Explanation: Strings "aaab" and "baa" are consistent since they only contain characters 'a' and 'b'.
//Example 2:
//
//Input: allowed = "abc", words = ["a","b","c","ab","ac","bc","abc"]
//Output: 7
//Explanation: All strings are consistent.
//Example 3:
//
//Input: allowed = "cad", words = ["cc","acd","b","ba","bac","bad","ac","d"]
//Output: 4
//Explanation: Strings "cc", "acd", "ac", and "d" are consistent.
///

type CountNumberOfConsistentLetters interface {
	SolutionA(allowed string, words []string) int
	Optimized(allowed string, words []string) int
}

type Impl struct{}

func NewCountNumberOfConsistentLetters() CountNumberOfConsistentLetters {
	return &Impl{}
}

func (i *Impl) SolutionA(allowed string, words []string) int {
	set := utils.NewStringHashSet()
	for _, char := range allowed {
		set.Add(string(char))
	}
	count := 0
	for _, word := range words {
		add := 1
		for _, char := range word {
			if ok := set.Contains(string(char)); !ok {
				add = 0
				break
			}
		}
		count += add
	}

	return count
}

func (i *Impl) Optimized(allowed string, words []string) int {
	allowedCharacters := 0
	for _, char := range allowed {
		allowedCharacters |= 1 << (int(char) - int('a'))
	}

	count := 0
	for _, word := range words {
		add := 1
		for _, char := range word {
			if (allowedCharacters & (1 << (int(char) - int('a')))) == 0 {
				add = 0
				break
			}
		}
		count += add
	}
	return count
}
