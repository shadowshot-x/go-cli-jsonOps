package jsonvalidator

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func CheckFP(pth string) (bool, error) {
	if _, err := os.Stat(pth); err == nil {
		return true, nil
	} else {
		return false, err
	}

}

func push(stack *[]string, val string) {
	*stack = append(*stack, val)
}

func printFileLineError(content string, index int) {
	cntr := -1
	for i := 0; i < len(content); i++ {

		if i == index {
			cntr = 1
		}
		if content[i] == 10 && cntr == 1 {
			fmt.Printf("<==== CHECK HERE")
			cntr = -1
		}
		fmt.Printf("%c", content[i])
	}
	if index == len(content)-1 {
		fmt.Printf("<==== CHECK HERE")
		cntr = -1
	}
	fmt.Println()
}

func ValidateJSON(content string, verbose int) bool {
	// we use slice to make stack
	var stack []string
	//pushing  => stack = append(stack, "{")
	//popping => stack = stack[:len(stack)-1]
	payload := strings.TrimSpace(content)
	lineNo := 1
	charCounter := -1
	for i := 0; i < len(payload); i++ {
		//checks for the current line number
		if payload[i] == 10 {
			lineNo = lineNo + 1
		}
		if payload[i] == '"' {
			if charCounter == -1 {
				charCounter = 1
			} else {
				charCounter = -1
			}
		}
		if (payload[i] >= 'a' && payload[i] <= 'z') || (payload[i] >= 'A' && payload[i] <= 'Z') {
			if charCounter == -1 {
				if verbose == 1 {
					fmt.Printf("DIRECTIONS :: Incorrect Usage of Double Quotes. Brackets Complete at Line No. %d\n", lineNo)
					printFileLineError(content, i)
				}
				return false
			}
		}
		if payload[i] == '{' {
			if i != 0 && len(stack) == 0 {
				if verbose == 1 {
					fmt.Printf("DIRECTIONS :: Incorrect Bracket Usage. Brackets Complete at Line No. %d\n", lineNo)
					printFileLineError(content, i)
				}

				return false
			}
			push(&stack, string(payload[i]))
		}
		if payload[i] == '}' {
			if len(stack) <= 0 {
				if verbose == 1 {
					fmt.Printf("DIRECTIONS :: Check Your Bracket at line No. %d in the text file\n", lineNo)
					printFileLineError(content, i)
				}
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) > 0 {
		printFileLineError(content, len(content)-1)
		return false
	}
	return true
}

func CLIExecuter() {
	fmt.Println("###### JSON VALIDATOR BY SHADOW ######")

	// declaration of flags
	// INT type flag for verbose
	// BOOL type flag for pretty
	verbosePtr := flag.Int("verbose", 1, "0 for no error description and 1 for error description")
	prettyPtr := flag.Int("pretty", 0, "Generate a pretty JSON file from text file")

	// Parse the Input from CLI to the ptrs declared above
	flag.Parse()

	// check for correct number of arguments
	if flag.NArg() != 1 && flag.NArg() != 2 {
		fmt.Println("Incorrect Number of Arguments Provided")
		os.Exit(1)
	}

	// 1. Check if the file provided exists.
	// 2. Read the File Provided into a string
	// 3. Write the Validation Logic
	// 4. Write the Pretty Logic
	// 5. Modify as per verbose
	// 6. Create the file OR Log the errors

	// checks if a file exists
	fileExist, err := CheckFP(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
	}
	if !fileExist {
		fmt.Println("Please Provide a valid file path")
		os.Exit(1)
	}

	// reading the file provided
	fileBytes, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
	}
	fileContent := string(fileBytes)
	// fmt.Println(fileContent)

	// Validate the File Content
	check := ValidateJSON(fileContent, *verbosePtr)
	if !check {
		fmt.Println("Your JSON is Incorrectly Formatted!!!!")
	} else {
		fmt.Println("Your JSON is Correctly Formatted!")
		//writing the logic for generating a pretty json file
		if *prettyPtr == 1 {
			err := ioutil.WriteFile("shadow-pretty.json", fileBytes, 0644)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Saved Your Pretty JSON file in ./shadow-pretty.json")
		}
	}

	os.Exit(0)
}
