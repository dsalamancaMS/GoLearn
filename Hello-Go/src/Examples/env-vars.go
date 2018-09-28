package main

import (
	"fmt"
	"os"
)

func main() {

	//	for _, env := range os.Environ() { Prints all Environments from System
	//		fmt.Println(env)
	//	}

	name := os.Getenv("USER")

	fmt.Println("Currently logged on user is", name)

}
