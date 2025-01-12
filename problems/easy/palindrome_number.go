// 9. Palindrome Number

// Given an integer x, return true if x is a palindrome, and false otherwise.

// Example 1:
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

// Example 2:
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

// Example 3:
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

// Constraints:
// -231 <= x <= 231 - 1

// Topics
// Math

package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := isPalindrome(121)
	b := isPalindrome(-121)
	c := isPalindrome(10)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	d := isPalindrome2(121)
	e := isPalindrome2(-121)
	f := isPalindrome2(10)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}

func isPalindrome(x int) bool {

	var reversed string
	xString := strconv.Itoa(x)

	for i := len(xString); i >= 1; i-- {

		reversed = fmt.Sprintf("%v%v", reversed, string(xString[i-1]))

	}

	if xString == reversed {
		return true
	}

	return false

}

func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedHalf := 0
	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}

	return x == reversedHalf || x == reversedHalf/10
}
