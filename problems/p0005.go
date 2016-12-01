/*
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import "fmt"

func gcd (a, b int) int {
	if b == 0 {
		return a;
	} else {
		return gcd(b, a%b);
	}
}

func lcm(n int) int {

	var multiple int = 1;

	for i:=2; i<=n; i++ {
		multiple *= i / gcd(i, multiple);
	}

	return multiple;
}

func main() {

	fmt.Print(lcm(20));

}