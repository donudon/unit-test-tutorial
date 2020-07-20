package main

import (
	"fmt"

	animal_prefix "github.com/donudon/unit-test-tutorial/animal-prefix"
)

func main(){
	animal, err := animal_prefix.GetAnimal("d")
	if err != nil {
		panic(err)
	}
	fmt.Println(animal)
}
