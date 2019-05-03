package cm

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

// GetStringFlag trying to read a flag value, if it does not exists ask an input from the user
func GetStringFlag(flagValue string, defaultValue string, text string) string {
	if len(flagValue) == 0 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(text)
		if len(defaultValue) > 0 {
			fmt.Print(" (" + defaultValue + "): ")
		} else {
			fmt.Print(": ")
		}
		answer, _ := reader.ReadString('\n')
		if len(answer) == 0 || answer == "\n" {
			if len(defaultValue) == 0 {
				fmt.Println("Input cannot be empty!")
				os.Exit(1)
			}
			answer = defaultValue
		}
		return strings.TrimSpace(answer)
	}
	return flagValue
}

// GetPassword trying to read a password flag value, if it does not exists ask an input from the user
func GetPassword(flagValue string, text string) string {
	if len(flagValue) == 0 {
		fmt.Print(text + ": ")
		if terminal.IsTerminal(0) {
			var fd = 0
			bytePassword, err := terminal.ReadPassword(fd)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			password := string(bytePassword)
			fmt.Println()
			return strings.TrimSpace(password)
		}
		answer, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if len(answer) == 0 || answer == "\n" {
			fmt.Println("Password cannot by empty")
			os.Exit(1)
		}
		return strings.TrimSpace(answer)

	}
	return flagValue
}

// EvaluateBoolValueFromString get a string boolean answer and evaluate as a boolean value
func EvaluateBoolValueFromString(answer string) bool {
	trueAnswerList := []string{"y", "yes", "true", "1"}
	result := false
	for _, v := range trueAnswerList {
		if v == strings.ToLower(answer) {
			result = true
		}
	}
	return result
}
