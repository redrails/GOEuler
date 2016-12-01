/*
By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/

package main

import "fmt"


func main() {

	sum := 0;
	a,b := 0,1;

	for b<4000000 {
		if b%2 == 0{
			sum+=b;
		}
		a, b = b, a+b;
	}

	fmt.Println(sum);

}