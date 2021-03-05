package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 0, 0}
	fmt.Println(threeSum(nums))
}

// 32ms
// 7MB
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	results := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			lo := i + 1
			hi := len(nums) - 1
			sum := -nums[i]
			for lo < hi {
				if nums[lo]+nums[hi] == sum {
					results = append(results, []int{nums[i], nums[lo], nums[hi]})
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					lo++
					hi--
				} else if nums[lo]+nums[hi] < sum {
					lo++
				} else {
					hi--
				}
			}
		}
	}
	return results
}
