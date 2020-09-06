package main

import (
	"log"

	"github.com/donudon/unit-test-tutorial/animal_prefix"
)

func main() {
	animal, err := animal_prefix.GetAnimal("d")
	if err != nil {
		log.Println("GetAnimal error: ", err)
	}
	log.Println(animal)
}
