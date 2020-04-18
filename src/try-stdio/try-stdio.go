package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var readText string
	readText = scanner.Text()
	fmt.Println(readText)

	//// Another version (Use Type inference)
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// readText := scanner.Text()
	// fmt.Println(readText)
}
