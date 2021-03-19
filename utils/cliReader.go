package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadNumber() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nEnter the number:")

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	number, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalf("your number is not a number: %s", err.Error())
	}

	return number
}
