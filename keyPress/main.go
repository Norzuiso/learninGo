package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	EscKey   = 27
	CtrlCKey = 3
	PipeKey = 124
)

func main() {
	fmt.Println("Welcome to ASCII Key Press")
	fmt.Println("Press ESC, Ctrl-C or | to finish")
	fmt.Println("Press any key follow by Enter to get the ASCII code")

	for {
		consoleReader := bufio.NewReaderSize(os.Stdin, 1)
		fmt.Print(">")
		input, _ := consoleReader.ReadByte()
		ascii := input
		if ascii == EscKey || ascii == CtrlCKey || ascii == PipeKey{
			fmt.Println("Thanks to use my program!")
			os.Exit(0)
		}
		fmt.Println("ASCII : ", ascii)
	}
}
