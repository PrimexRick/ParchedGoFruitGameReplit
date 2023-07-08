package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	X := "cucumber"
	Y := "banana"
	Z := "watermelon"
	A := "pineapple"
	B := "jackfruit"

	fmt.Println("Welcome! Enter a command (1-5) or 'exit' to quit.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Command: ")
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}

		h, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number or 'exit' to quit.")
			continue
		}

		switch h {
		case 1:
			fmt.Println("Put", X, "in your ass.")
		case 2:
			fmt.Println("Put", Y, "in your ass.")
		case 3:
			fmt.Println("Put", Z, "in your ass.")
		case 4:
			fmt.Println("Put", A, "in your ass.")
		case 5:
			fmt.Println("Put", B, "in your ass.")
		default:
			fmt.Println("Invalid command. Ask Your StepDad To Shove His Dick In Your Mouth.")
		}
	}
}
