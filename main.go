package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// credential loaded from environment variable
var apiKey = os.Getenv("API_KEY")

func main() {
	fmt.Println("Hello, World!")

	fmt.Print("Enter a number: ")
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && line == "" {
		fmt.Println("Invalid input: please enter a valid integer.")
		return
	}
	number, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		fmt.Println("Invalid input: please enter a valid integer.")
		return
	}
	fmt.Println(number * 2)
}
