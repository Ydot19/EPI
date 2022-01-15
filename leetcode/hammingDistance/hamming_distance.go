package hammingDistance

//The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
//
//Given two integers x and y, return the Hamming distance between them.
//
//
//
//Example 1:
//
//Input: x = 1, y = 4
//Output: 2
//Explanation:
//1   (0 0 0 1)
//4   (0 1 0 0)
//       ↑   ↑
//The above arrows point to positions where the corresponding bits are different.
//Example 2:
//
//Input: x = 3, y = 1
//Output: 1
//
//
//Constraints:
//
//0 <= x, y <= 231 - 1
//

type HammingDistance interface {
	Optimized(int, int) int
}

type Impl struct {
}

func NewHammingDistance() HammingDistance {
	return &Impl{}
}

func (i *Impl) Optimized(x int, y int) int {
	dist := 0
	if x == y {
		return dist
	}
	xor := x ^ y
	for xor > 0 {
		dist += 1
		xor &= xor - 1 // remove the least significant bit
	}
	return dist
}
