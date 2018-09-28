package main

import (
	"fmt"
	"os"
)

func main() {

	ifFunction()
	switchFunc("docker")
	errorHandling()

}

/* you can declare variables and other stuff before evaluating boolean,
they get garbage collected, ONLY BOOLEAN on
if <simple statement>;<boolean evaluation>*/

func ifFunction() {

	if firstRank, secondRank := 39, 25; firstRank < secondRank {
		fmt.Println("\nThe First is better than the second")
	} else if firstRank > secondRank {
		fmt.Println("\nThe Second one is doing better")
	} else {
		fmt.Println("\nBoth are equal")
	}

}

func switchFunc(option string) {

	switch option {

	case "linux":
		fmt.Println("\nHello Linux")
	case "docker":
		fmt.Println("\nHello Docker") //fallthrough can be used as break to use the next case
	case "kubernetes":
		fmt.Println("\nKubectl hello")
	default:

	}
}

func errorHandling() error {

	_, err := os.Open("/tmp/hello")
	if err != nil {
		fmt.Println("ERROR! ", err)
	}
	return err
}
