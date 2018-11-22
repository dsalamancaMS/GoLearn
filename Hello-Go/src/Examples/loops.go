package main

import (
	"fmt"
	"time"
)

func main() {
	//	theFor()
	theRangeFor()
}

/* for <expression> "expression" will always be true if its not evaluated

for i := range thislist --> "range list" is an array. the "i" will be each item on the list

for i :=0; i < 10; i++ ---> the first statement runs only at the begining, the second is the boolean eval, the last will run on each iteration

*/

func theFor() {

	for timer := 5; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("BOOM!")
			break
		} else {
			fmt.Println(timer)
		}

		time.Sleep(1 * time.Second)
	}

}

func theRangeFor() {
	names := []string{"Denis", "John",
		"Mike"}
	for x, i := range names {
		fmt.Println("Name", i, "is on", x)
	}
}
