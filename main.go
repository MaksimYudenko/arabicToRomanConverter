package main

import (
	"fmt"
	"log"

	"github.com/MaksimYudenko/andersen/service"
	"github.com/MaksimYudenko/andersen/utils"
)

func main() {
	input := utils.ReadNumber()

	output, err := service.TransformArabicToRoman(input)
	if err != nil {
		log.Fatalf("check an input number: %s", err.Error())
	}

	fmt.Printf("\nNumber %d is represented as %v\n", input, output)
}
