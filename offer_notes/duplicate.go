/**
 * description: 在一个长度为 n 的数组里的所有数字都在 0 到 n-1 
 * 的范围内。数组中某些数字是重复的，但不知道有几个数字是重复的。也不知道每个数字重复几次。请找出数组中任意一个重复的数字。例如，如果输入长度为 7 
 * 的数组 {2, 3, 1, 0, 2, 5, 3}，那么对应的输出是第一个重复的数字 2。
 * 要求复杂度为 O(N) + O(1)，时间复杂度 O(N)，空间复杂度 O(1)。因此不能使用排序的方法，也不能使用额外的标记数组。
 */
package main

import "fmt"

func main() {
	arr := []int{2, 3, 1, 0, 2, 5}
	fmt.Print(duplicate(arr))
}

// [2, 3, 1, 0, 2, 5] arr
func duplicate(nums []int) bool {
	if nums == nil ||  len(nums) == 0 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println("nums:",nums)
		for (nums[i] != i && nums[nums[i]] != nums[i]) {  // for 不是if 
			fmt.Println("swap:",nums[i], i, nums)
			swap(nums[i], i, nums)
		}

		if (nums[i] != i && nums[nums[i]] == nums[i]) {
			fmt.Println("Result: Position->", i, " Value->", nums[i])

			return true
		}
	}
	return false
}

func swap(i int, j int, nums []int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t	
}