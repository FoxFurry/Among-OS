package command

import (
	"fmt"
	"os"
)

type shutdown struct {}

func (e *shutdown) Execute(parameters []string) string{
	if len(parameters) != 0 {
		return "Not expected any parameters for <$ shutdown>"
	}
	fmt.Println("Shutting down Among OS")
	os.Exit(0)
	return ""
}

func (e *shutdown) Help() string {
	return "shutdown closes Among OS"
}
