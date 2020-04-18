package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func myAdd(a int, b int) int {
	return a + b
}

func main() {
	var intA int
	var intB int
	var err error
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(os.Stdin)

	fmt.Println("Input intA:")
	scanner.Scan()
	intA, err = strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Cannot convert to int.")
		return
	}

	fmt.Println("Input intB:")
	scanner.Scan()
	intB, err = strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Cannot convert to int.")
		return
	}

	fmt.Println("Result:")
	fmt.Println(myAdd(intA, intB))
}
