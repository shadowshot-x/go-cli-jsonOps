package main

import (
	"github.com/abiosoft/ishell/v2"
	"knowsearch.ml/jsondigger"
	"knowsearch.ml/jsonvalidator"
)

func main() {
	// jsonvalidator.CLIExecuter()
	shell := ishell.New()
	shell.AddCmd(&ishell.Cmd{
		Name: "jsonvalidator",
		Help: "Validate Your JSON file and Find location of an error if present.",
		Func: func(c *ishell.Context) {
			jsonvalidator.CLIExecuter()
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "jsondigger",
		Help: "Get the Values of your JSON file and Query the JSON file as objects in Real Time",
		Func: func(c *ishell.Context) {
			jsondigger.CLIExecuter(c)
		},
	})

	shell.Run()
}
