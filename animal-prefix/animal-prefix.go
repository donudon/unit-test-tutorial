package animal_prefix

import (
	"errors"
	"strings"
)

func GetAnimal(firstLetter string) (animal string, err error){
	if strings.ToLower(firstLetter) == "d" {
		animal = "Dog"
	} else if strings.ToLower(firstLetter) == "c" {
		animal = "Cat"
	} else {
		err = errors.New("Undefined Animal")
	}
	return
}
