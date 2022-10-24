package prompt

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
	"syscall"
)

func Prompt(text string, defaultValue string) string {
	fmt.Println(text)
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, " \n")

	if input == "" {
		return defaultValue
	}

	return input
}

func YesNo(text string, defaultValue bool) bool {
	yn := fmt.Sprintf(" [y/%s]", bold("n"))
	if defaultValue {
		yn = fmt.Sprintf(" [%s/n]", bold("y"))
	}

	fmt.Print(text + " " + yn + ": ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, " \n")

	if input == "" {
		return defaultValue
	}

	return input == "y"
}

func Password() string {
	fmt.Print("Password: ")
	passwordAsBytes, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return ""
	}

	fmt.Println("")
	return string(passwordAsBytes)
}

func bold(input string) string {
	return fmt.Sprintf("\u001B[1m%s\u001B[0m", input)
}
