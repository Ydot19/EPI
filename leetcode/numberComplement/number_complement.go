package numberComplement

import "math/bits"

//The complement of an integer is the integer you get when you flip all the 0's to 1's and all the 1's to 0's in its binary representation.
//
//For example, The integer 5 is "101" in binary and its complement is "010" which is the integer 2.
//Given an integer num, return its complement.
//
//
//
//Example 1:
//
//Input: num = 5
//Output: 2
//Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.
//Example 2:
//
//Input: num = 1
//Output: 0
//Explanation: The binary representation of 1 is 1 (no leading zero bits), and its complement is 0. So you need to output 0.

type NumberComplement interface {
	SolutionA(num int) int
	Optimized(num int) int
}

type Impl struct {
}

func NewNumberComplement() NumberComplement {
	return &Impl{}
}

// Optimized explained.
// Find the bit length (uint(num)) and use XOr (at least one difference at the bit position)
func (i *Impl) Optimized(num int) int {
	return (1<<(bits.Len(uint(num))) - 1) ^ num
}

// SolutionA explained.
// Flip bits one by one.
func (i *Impl) SolutionA(num int) int {
	bitIndex := 1
	ret := num

	for ret >= bitIndex {
		ret ^= bitIndex
		bitIndex <<= 1
	}
	return ret
}
