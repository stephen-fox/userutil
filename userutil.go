package userutil

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

// PromptOptions configures how the input prompt should behave.
type PromptOptions struct {
	// Whether or not to hide the user's input.
	ShouldHideInput bool

	// The message appended to Y/N prompts.
	YesNoMessage string

	// The string to prefix to the user input message prompt.
	InputPrefix string

	// The string to append to the user input message prompt.
	InputSuffix string
}

func (o PromptOptions) YesNoSuffixFormat() string {
	if len(o.YesNoMessage) == 0 {
		return " [Y/N]"
	}

	return o.YesNoMessage
}

func (o PromptOptions) InputPrefixFormat() string {
	if len(o.InputPrefix) == 0 {
		return "> "
	}

	return o.InputPrefix
}

func (o PromptOptions) InputSuffixFormat() string {
	if len(o.InputSuffix) == 0 {
		return ": "
	}

	return o.InputSuffix
}

// GetYesOrNoUserInput issues a prompt and then checks if the user answered
// 'yes', 'y', 'no', or 'n'. Any capitalization of these inputs is valid. If
// the user did not provide any of these inputs then an error of type
// InputError is returned. If the user answered in the affirmative, then true
// is returned.
func GetYesOrNoUserInput(promptMessage string, options PromptOptions) (bool, error) {
	promptMessage = promptMessage + options.YesNoSuffixFormat()

	result, err := GetUserInput(promptMessage, options)
	if err != nil {
		return false, err
	}

	result = strings.ToLower(result)
	switch result {
	case "yes":
		fallthrough
	case "y":
		return true, nil
	case "no":
		fallthrough
	case "n":
		return false, nil
	}

	return false, InputError{
		reason:         "Please specify 'y', 'yes', 'n', or 'no'",
		didNotUseYesNo: true,
	}
}

// GetUserInput issues a prompt to the user and records the user's response.
func GetUserInput(promptMessage string, options PromptOptions) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	if len(strings.TrimSpace(promptMessage)) > 0 {
		fmt.Print(options.InputPrefixFormat() + promptMessage + options.InputSuffixFormat())
	}

	input := ""
	if options.ShouldHideInput {
		state, err := terminal.MakeRaw(0)
		if err != nil {
			return "", err
		}

		// Restore the terminal state if the program is Control+C'ed.
		onInterrupts := make(chan os.Signal, 1)
		signal.Notify(onInterrupts, os.Interrupt)

		defer signal.Stop(onInterrupts)
		defer close(onInterrupts)

		go func() {
			for range onInterrupts {
				terminal.Restore(0, state)
				fmt.Println()
				os.Exit(0)
			}
		}()

		raw, err := terminal.ReadPassword(0)
		if err != nil {
			return "", err
		}

		terminal.Restore(0, state)
		fmt.Println()

		input = string(raw)
	} else {
		readInDelimiter := '\n'
		var err error

		input, err = reader.ReadString(byte(readInDelimiter))
		if err != nil {
			return "", err
		}

		input = strings.TrimSuffix(input, string(readInDelimiter))
	}

	return input, nil
}
