/*
Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import "fmt"
import "strconv"

func main() {
	var s string = "";
	var n,p,z int = 0,0,0;

	for a := 999; a > 800; a-- {
		L: for b := 999; b > 800; b-- {
			n = a * b
			s = strconv.Itoa(n)
			z = int(len(s)/2)
			for i := 0; i < z; i++ {
				if s[i] != s[len(s) - i - 1] { continue L }
			}
			if p == 0 || n > p { p = n }
		}
	} 
}