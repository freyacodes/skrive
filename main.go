package main

import (
	"fmt"
	"os"

	"skrive/data"
	"skrive/data/fs"
	"skrive/log"
	"skrive/logic"
	"skrive/startMenu"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	if parseErr := parse(); parseErr != nil {
		printHelp(parseErr)
		os.Exit(1)
	} else if *helpFlag {
		printHelp(nil)
		os.Exit(0)
	}

	err := logic.Setup(*fileArg)

	if err == nil && subcommand != nil {
		handleSubcommands()
	}

	if err == nil {
		var model tea.Model
		if subcommand != nil && *subcommand == "log" {
			model, _ = log.InitializeModel(func() (tea.Model, tea.Cmd) {
				return model, tea.Quit
			})
		} else {
			model = startMenu.InitializeModel()
		}

		_, err = tea.
			NewProgram(model).
			Run()
	}

	exitIfError(err)
}

func initialiseStorageInterface() data.Storage {
	// There is currently only one type of storage
	p, err := fs.GetPath(*fileArg)
	exitIfError(err)
	return fs.FsStorage{Path: *p}
}

func handleSubcommands() {
	switch *subcommand {
	case "log":
		if len(positionalArguments) == 0 {
			// Handled in Bubbletea initialization code
			return
		}
		var storage = initialiseStorageInterface()
		exitIfError(log.Invoke(storage, positionalArguments))
	}
	os.Exit(0)
}

func exitIfError(err error) {
	if err == nil {
		return
	}
	fmt.Println("Undskyld! Something went wrong >w< here it is: %v", err)
	os.Exit(1)
}
