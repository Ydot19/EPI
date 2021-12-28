package numberofStepsToReduceToZero

///
//Given an integer num, return the number of steps to reduce it to zero.
//
//In one step, if the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.
//
//
//
//Example 1:
//
//Input: num = 14
//Output: 6
//Explanation:
//Step 1) 14 is even; divide by 2 and obtain 7.
//Step 2) 7 is odd; subtract 1 and obtain 6.
//Step 3) 6 is even; divide by 2 and obtain 3.
//Step 4) 3 is odd; subtract 1 and obtain 2.
//Step 5) 2 is even; divide by 2 and obtain 1.
//Step 6) 1 is odd; subtract 1 and obtain 0.
//Example 2:
//
//Input: num = 8
//Output: 4
//Explanation:
//Step 1) 8 is even; divide by 2 and obtain 4.
//Step 2) 4 is even; divide by 2 and obtain 2.
//Step 3) 2 is even; divide by 2 and obtain 1.
//Step 4) 1 is odd; subtract 1 and obtain 0.
//Example 3:
//
//Input: num = 123
//Output: 12
//
//
//Constraints:
//
//0 <= num <= 106
///

type NumberOfSteps interface {
	SolutionA(int) int
	Optimized(int) int
}

type Impl struct{}

func NewNumberOfStepsImpl() NumberOfSteps {
	return &Impl{}
}

func (n *Impl) SolutionA(num int) int {
	if num == 0 {
		return num
	}
	steps := 0

	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		steps += 1
	}
	return steps
}

func (n *Impl) Optimized(num int) int {
	if num == 0 {
		return num
	}

	// see: https://codeforwin.org/2018/05/10-cool-bitwise-operator-hacks-and-tricks.html
	var steps int // without defining sets to 0 value
	for num > 0 {
		// if odd, add 1 then add one for an extra addition
		// num & 1 evaluates to 1 for odd numbers, and zero for even
		steps += num&1 + 1
		num >>= 1 // shift right, equivalent to floor division
	}
	return steps - 1
}
