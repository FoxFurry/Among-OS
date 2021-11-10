package command

import (
	"fmt"
)

type nothing struct{}

func (e *nothing) Execute(parameters []string) string {
	if len(parameters) != 0 {
		return "Not expected any parameters for <$ shutdown>"
	}
	fmt.Println(" ")
	return ""
}

func (e *nothing) Help() string {
	return "nothing happened"
}
