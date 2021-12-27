package twosum

type TwoSum interface {
	Optimized(nums []int, target int) []int
}

type Impl struct {
}

func NewTwoSum() *Impl {
	return &Impl{}
}

//func (ts *TwoSumImpl) BruteForce(nums []int, target int) []int {
//	return []int{}
//}

func (ts *Impl) Optimized(nums []int, target int) []int {
	alreadySean := make(map[int]int)
	for i, val := range nums {
		diff := target - val
		if j, ok := alreadySean[diff]; ok {
			return []int{i, j}
		} else {
			alreadySean[val] = i
		}
	}
	return []int{}
}
