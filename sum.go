//program to print sum of all numbers below 50 completely divisible by 3 or 5

package main

import(
      "fmt"
     )

func main() {
	sum := 0 
	for i := 1; i < 50; i++{
      		if i % 3 == 0 {
 			sum = sum + i
		}else {
                       if i % 5 ==0 {
			sum = sum + i
		      }
                }
             }
       fmt.Println(sum)
}

