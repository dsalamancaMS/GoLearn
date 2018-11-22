package main

import (
	"fmt"
)

func main() {
	//createMapByMake()
	//createMapByCompositeForm()
	printMyBankAccunts()
}

func createMapByMake() {
	bankAccount := make(map[string]int)
	bankAccount["Dolar Account"] = 500
	bankAccount["Euros Account"] = 300
	fmt.Printf("\n Your Bank account: %v \n", bankAccount)
}

func createMapByCompositeForm() {

	bankAccount := map[string]int{
		"Yuan": 3000,
		"Yens": 5000,
	}

	fmt.Printf("\n Your International Bank accounts: %v\n", bankAccount)
}

func printMyBankAccunts() {

	bankAccount := map[string]int{
		"Yuans":  3000,
		"Yens":   5000,
		"Dolars": 100,
		"Euros":  12200,
	}
	for key, value := range bankAccount {
		fmt.Printf("My Bank Account in %v has %v remaining\n", key, value)
	}

	/// to delete an item:

	delete(bankAccount, "Yuans")
}
