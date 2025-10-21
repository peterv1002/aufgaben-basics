package lists

import "sort"

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	return nums[0]

}
