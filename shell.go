package main

import (
	"bufio"
	"fmt"
	"os"
)

func containsString(slice []string, value string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			return true
		}
	}
	return false
}

func parseCommand(str string) []string {
	var output []string = []string{""}

	for i := 0; i < len(str); i++ { // Parse the command by symbols
		if string(str[i]) == " " {
			output = append(output, "")
		} else {
			output[len(output)-1] += string(str[i])
		}
	}

	count := 0

	// Delete empty elements from the output slice in case of multiple spaces between command arguments.

	for count != (len(str)-1)/2 {
		for i := 0; i < len(output); i++ {
			if output[i] == "" {
				output = append(output[:i], output[i+1:]...)
			}
		}
		count++
	}
	output[len(output)-1] = output[len(output)-1][:len(output[len(output)-1])-2]
	return output
}

func executeCommand(args []string) (output string) {
	if args[0] == "encode" {
		if len(args) == 1 {
			return "Not enough arguments. Need a string value to encode"
		} else {
			if containsString(args, "--full-string") {
				return encodeToFullString(args[1])
			} else {
				output := ""
				for i := 0; i < len(encodeLehaCode(args[1])); i++ {
					output += "["
					output += encodeLehaCode(args[1])[i].value
					output += "] "
				}
				return output
			}
		}

	} else if args[0] == "decode" {
		if len(args) == 1 {
			return "Not enough arguments. Need a string value to decode"
		} else {
			return decodeLehaCode(args[1])
		}
	} else {
		return "Command '" + args[0] + "' doesn't exist"
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(">")
		input, _ := reader.ReadString('\n')
		fmt.Println(executeCommand(parseCommand(input)))
	}

}
