package epi

func main() {

	a := func(nums []int, target int) []int{
		already_sean := make(map[int]int)
		for i, val := range nums {
			diff := target - val
			if j, ok := already_sean[diff]; ok {
				return []int{i, j}
			} else {
				already_sean[val] = i
			}
		}
		return []int{}

	}
}