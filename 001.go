func twoSum(nums []int, target int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]bool)
	for i, num := range nums {
		m1[num] = i
		m2[num] = true
	}
	for i, num := range nums {
		if m2[target-num] && m1[target-num] != i {
			return []int{i, m1[target-num]}
		}
	}
	return []int{}
}
