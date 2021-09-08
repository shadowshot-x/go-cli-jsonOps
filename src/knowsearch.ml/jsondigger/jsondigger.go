package jsondigger

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/abiosoft/ishell/v2"
	"github.com/fatih/color"
)

func DisplayObject(mp map[string]interface{}, c *ishell.Context, lev int) {
	// define the colours
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintfFunc()

	// go over the key value pairs in the map
	for key, ele := range mp {
		// here we check if the current ele is of type map string interface
		_, ok := ele.(map[string]interface{})
		// For indentation
		for i := 0; i < lev; i++ {
			c.Printf("\t")
		}
		// If not of map string interface type, print the pair
		if !ok {
			c.Println(cyan(key), " : ", ele)
		} else {
			// go down the level by calling DisplayObject func again
			// here lev is used to keep track of indentation
			c.Println(yellow("Enter the Map of ", key))
			DisplayObject(ele.(map[string]interface{}), c, lev+1)
			for i := 0; i < lev; i++ {
				c.Printf("\t")
			}
			c.Println(yellow("Exiting the Map of ", key))
		}
	}
}

func QueryJSON(mp map[string]interface{}, c *ishell.Context, splitQuery []string, i int) error {
	// check for out of bounds
	if i >= len(splitQuery) {
		return errors.New("please enter a valid query")
	}
	// get the current value for the current key
	_, ok := mp[splitQuery[i]]
	if !ok {
		// if this value does not exist, we can be sure that this was a wrong query.
		c.Println("Incorrect Query Provided")
		return errors.New("please enter a valid query")
	}
	// if this exists and it is the last element of the array, we print the string and return
	if i == len(splitQuery)-1 {
		c.Println(mp[splitQuery[i]])
		return nil
	} else {
		// else we keep on querying to the next level
		return QueryJSON(mp[splitQuery[i]].(map[string]interface{}), c, splitQuery, i+1)
	}
}

// element for the Stack.
type StackEle struct {
	name string
	util map[string]interface{}
}

func CLIExecuter(c *ishell.Context) {
	c.ShowPrompt(false)

	// take user input of the file provided
	c.Print("Enter the File Path\n")
	filepath := c.ReadLine()

	// reading the file provided
	file, err := os.Open(filepath)
	if err != nil {
		// return with an error indication
		c.Print(err)
		c.ShowPrompt(true)
		os.Exit(1)
	}
	defer file.Close()

	// Now, we want to read the file line by line. So,I have used Bufio over here.
	scanner := bufio.NewScanner(file)
	jsonContent := make(map[string]interface{})

	// make a stack of type of StackEle with string and util.
	var stack []StackEle
	// append a base element
	stack = append(stack, StackEle{"start", jsonContent})

	scanner.Scan()
	// now we read the JSON line by line
	for scanner.Scan() {
		// store the current text in a string and deal with whitespace
		sc := scanner.Text()
		sc = strings.TrimSpace(sc)
		splitString := strings.Split(sc, ":")
		key := strings.Trim(splitString[0], "\"")

		// This is the case where nesting begins
		if len(splitString) == 2 && strings.TrimSpace(splitString[1])[0] == '{' {
			// add the key value pair to the current stack's tops
			stack[len(stack)-1].util[key] = make(map[string]interface{})
			// add the appended element to the stack
			stack = append(stack, StackEle{key, stack[len(stack)-1].util[key].(map[string]interface{})})

		} else if strings.TrimSpace(key)[0] == '}' {
			// this means we are coming down a level
			stack = stack[:len(stack)-1]

		} else if len(splitString) == 2 && strings.TrimSpace(splitString[1])[0] == '[' {
			//Deal with array
			substr := strings.TrimSpace(splitString[1])[1 : len(strings.TrimSpace(splitString[1]))-1]
			splits := strings.Split(strings.Trim(substr, "\""), ",")
			stack[len(stack)-1].util[key] = splits

		} else if len(splitString) == 2 && strings.TrimSpace(splitString[1])[0] == '"' {
			// deal with string
			stack[len(stack)-1].util[key] = splitString[1]

		} else {
			// this will mean that the JSON is not formatted as per our code
			c.Println("JSON is not formatted as per the Current Package Implementation")
			c.ShowPrompt(true)
			os.Exit(1)

		}
	}
	c.Println("Generating the Content Object of JSON")
	DisplayObject(jsonContent, c, 0)

	c.Println("Write your Query. To Stop, Simply Write 'N' and Click Enter")

	query := ""
	for query != "N" {
		query = c.ReadLine()
		if query == "N" {
			break
		}
		splitQuery := strings.Split(query, ".")
		err = QueryJSON(jsonContent, c, splitQuery, 0)
		if err != nil {
			c.Println(err)
		}
	}

	c.ShowPrompt(true)
}
