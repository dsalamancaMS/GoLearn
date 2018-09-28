package main

/*I AM A BLOCK COMMENT*/

//COMMENT

import (
	"fmt"
	"reflect" //reflect to check var type
	"runtime" //Println function
)

var (
	name   string
	module int
	course string
)

func main() {
	denis := 24.2     //dynamic var form not to use on package body only functions
	pointer := &denis // create a Var poining to the memory address
	fmt.Println("Hello from", runtime.GOOS)
	fmt.Println("\nMy name is", name)
	fmt.Println("\nCurrent Module is", module, "is type", reflect.TypeOf(module))
	fmt.Println("\nDenis is", denis, "and is type", reflect.TypeOf(denis))
	fmt.Println("Memory Address of 'denis' is", pointer) //reference memory address of a variable
	fmt.Println("And the value of 'denis' is", *pointer) //print contents of pointer
}
