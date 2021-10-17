package animal_prefix

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrCode1 = errors.New("Undefined Animal")
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
	if len(animal) == 0 {
		err = ErrCode1
		return
	}
	letter = fmt.Sprintf(string([]rune(animal)[0]))
	return
}
