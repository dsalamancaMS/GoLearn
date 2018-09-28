package main

import (
	"fmt"
	"strings"
)

func main() {

	module := "function basics"
	author := "denis salamanca"

	bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 16)

	fmt.Println(converter(module, author))

	fmt.Println(bestFinish)

}

func converter(s1, s2 string) (str1, str2 string) {
	s1 = strings.Title(s1)
	s2 = strings.ToUpper(s2)

	return s1, s2

}

func bestLeagueFinishes(finishes ...int) int { // ""..."before var type to indicate infinite number of parameters

	best := finishes[0]

	for _, i := range finishes {
		if i < best {
			best = i
		}
	}

	return best
}
