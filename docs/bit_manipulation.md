# Bit Manipulation

## Exclusive OR

- One of two bits must be 1 otherwise returns 0

Problems
- [Decoded XORed Array](../leetcode/decodedXORedArray)

## Bitwise And (&)

```text
a & n     # Bitwise and
```
- ```a & 1``` returns 1 if odd 0 if even, same as write `a % 2` (a modulo 2)

Problems
- [Number of Steps to Reduce To Zero](../leetcode/numberofStepsToReduceToZero)
- [Hamming Distance](../leetcode/hammingDistance)
- [Counting Bits](../leetcode/countingBits)

## Shift Right

```text
a >> n     # Bitwise right shift
```
- Same as floor division

Problems
- [Number of Steps to Reduce To Zero](../leetcode/numberofStepsToReduceToZero)


## Shift Left

```text
a << n     # Bitwise left shift
```
- Same as multiplication

Problems
- [Count Number of Consistent Strings](../leetcode/countNumberOfConsistentStrings)


## Inclusive OR (in-place bitwise OR operator)

```text
a |= 1 
```

- Same as +=
- In python 3.9+, can be used to merge two dictionaries

Problems:
- [Count Number of Consistent Strings](../leetcode/countNumberOfConsistentStrings)

## Special Scenarios

### Remove Least Significant B

```text
leastSigBit = leastSigBit & (leastSigBit - 1)
```


code example: 

```go
package example

// xor     = 100 = 0b1100100
// xor - 1 = 99  = Ob1100011
// & 1
//           96  = 0b1100000
//           95  = 0b1011111
// & 2
//           64  = 0b1000000
//           63  =  0b111111
// & 3
//               = 0b0
///
func ex(){
	xor := 100
	for xor > 0 {
		xor &= xor - 1
	}
}
```

Problems
- [Hamming Distance](../leetcode/hammingDistance)
