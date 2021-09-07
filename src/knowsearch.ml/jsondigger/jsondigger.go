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
	yellow := color.New(color.FgYellow).SprintFunc()
	// blue := color.New(color.BgHiBlue).SprintFunc()
	cyan := color.New(color.FgCyan).SprintfFunc()
	for key, ele := range mp {
		_, ok := ele.(map[string]interface{})
		for i := 0; i < lev; i++ {
			c.Printf("\t")
		}
		if !ok {
			c.Println(cyan(key), " : ", ele)
		} else {
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
	if i >= len(splitQuery) {
		return errors.New("please enter a valid query")
	}
	_, ok := mp[splitQuery[i]]
	if !ok {
		c.Println("Incorrect Query Provided")
		return errors.New("please enter a valid query")
	}
	if i == len(splitQuery)-1 {
		c.Println(mp[splitQuery[i]])
		return nil
	} else {
		return QueryJSON(mp[splitQuery[i]].(map[string]interface{}), c, splitQuery, i+1)
	}
}

type StackEle struct {
	name string
	util map[string]interface{}
}

func CLIExecuter(c *ishell.Context) {
	c.ShowPrompt(false)

	c.Print("Enter the File Path\n")
	filepath := c.ReadLine()

	// reading the file provided
	file, err := os.Open(filepath)
	if err != nil {
		c.Print(err)
		c.ShowPrompt(true)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	jsonContent := make(map[string]interface{})

	var stack []StackEle
	stack = append(stack, StackEle{"start", jsonContent})

	scanner.Scan()
	for scanner.Scan() {
		sc := scanner.Text()
		sc = strings.TrimSpace(sc)
		splitString := strings.Split(sc, ":")
		key := strings.Trim(splitString[0], "\"")
		// c.Printf("Line is    :%c\n", strings.TrimSpace(splitString[0])[0])
		// value := strings.Trim(splitString[1],"\"")
		if len(splitString) == 2 && strings.TrimSpace(splitString[1])[0] == '{' {
			stack[len(stack)-1].util[key] = make(map[string]interface{})
			stack = append(stack, StackEle{key, stack[len(stack)-1].util[key].(map[string]interface{})})
		} else if strings.TrimSpace(key)[0] == '}' {
			stack = stack[:len(stack)-1]
		} else if len(splitString) == 2 && strings.TrimSpace(splitString[1])[0] == '[' {
			substr := strings.TrimSpace(splitString[1])[1 : len(strings.TrimSpace(splitString[1]))-1]
			splits := strings.Split(strings.Trim(substr, "\""), ",")
			stack[len(stack)-1].util[key] = splits
		} else if len(splitString) == 2 && strings.TrimSpace(splitString[1])[0] == '"' {
			stack[len(stack)-1].util[key] = splitString[1]
		} else {
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
