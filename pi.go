// pi/4 = 1-1/3+1/5-1/7 ......
package main

import (
	"fmt"
)

func main() {
	fmt.Println("CPU !!!Daje Ve!!!")
	fmt.Println("input the times")
	a := 0
	fmt.Scan(&a)
	x := 0.0
	fenmu := 1.0

	for i := 0; i < a; i++ {
		if i%2 != 0 {
			x = x - 1/fenmu
		} else {
			x = x + 1/fenmu
		}
		fenmu += 2
	}

	pi := 0.0
	pi = x * 4
	fmt.Println(pi)
}
