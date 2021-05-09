package euler

import "strconv"
import "fmt"

/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
 */

func prob4(min, max int) int {
	var maxp int

	for n1 := min; n1 <= max; n1++{

		for n2 := min; n2 <= max; n2++{
			np := n1*n2
			if isPalindromic(np){
				if np > maxp {
					maxp = np
				}
			}

		}
	}
	fmt.Print ("maxp: ", maxp)
	return maxp
}

func isPalindromic(n int) bool {
	sn := strconv.Itoa(n)
	sr := []rune(sn)
	cnt := 0
	for i := len(sr) - 1; i >= 0; i-- {
		if cnt < i {
			if sr[i] != sr[cnt] {
				return false
			}
		}
		cnt++
	}
	return true
}
