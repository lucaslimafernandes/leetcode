// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]

// Constraints:

// 2 <= nums.length <= 10^4
// -109 <= nums[i] <= 10^9
// -109 <= target <= 10^9
// Only one valid answer exists.

// Topics
// Array
// Hash Table

package main

import "fmt"

func main() {

	a := twoSum([]int{2, 7, 11, 15}, 9)
	b := twoSum([]int{3, 2, 4}, 6)
	c := twoSum([]int{3, 3}, 6)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i, value := range nums {

		compl := target - value

		index, ok := m[compl]
		if ok {
			return []int{index, i}
		}

		m[value] = i

	}

	return nil

}
