package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if len(input) <= 1 {
			continue
		}
		// Handle the execution of the input.
		err = execInput(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// ErrNoPath is returned when 'cd' was called without a second argument.
var ErrNoPath = errors.New("path required")

func execInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split to get parallel commands.
	parArgs := strings.Split(input, "&")

	var trimmedArgs []string
	for _, arg := range parArgs {
		newArg := strings.TrimPrefix(arg, " ")
		trimmedArgs = append(trimmedArgs, strings.TrimSuffix(newArg," "))
	}

	//for _, arg := range trimmedArgs {
	//	fmt.Print(arg)
	//	fmt.Print("\n")
	//}
	// Split the input to separate the command and the arguments.

	for _, arg := range trimmedArgs[1:] {
		args := strings.Split(arg, " ")
		go execCommand(args)
		//if err != nil {
		//	return err
		//}
	}
	err := execCommand(strings.Split(trimmedArgs[0], " "))
	if err != nil {
		return err
	}
	//args := strings.Split(input, " ")
	return nil
}

func execCommand(args []string) error {
	// Pass the program and the arguments separately.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and save it's output.
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}