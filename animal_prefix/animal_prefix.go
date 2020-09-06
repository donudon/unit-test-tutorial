package animal_prefix

import (
	"errors"
	"fmt"
	"strings"
)

func GetAnimal(firstLetter string) (animal string, err error) {
	if strings.ToLower(firstLetter) == "d" {
		animal = "Dog"
	} else if strings.ToLower(firstLetter) == "c" {
		animal = "Cat"
	} else {
		err = errors.New("Undefined Animal Prefix")
	}
	return
}

func GetFirstLetterAnimal(animal string) (letter string, err error) {
	if len(animal) == 0{
		err = errors.New("Undefined Animal")
	}
	letter = fmt.Sprintf(string([]rune(animal)[0]))
	return
}

