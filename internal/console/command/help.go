package command

import "fmt"

type help struct {}

func (e *help) Execute(parameters []string) string{
	if len(parameters) == 0 {
		return e.Help()
	}

	allCommands := GetCommands()

	cmd, ok := allCommands[parameters[0]]
	if !ok {
		return fmt.Sprintf("command %s not found", parameters[0])
	}else {
		return cmd.Help()
	}
}

func (e *help) Help() string {
	return "Help is a simple command which show description of another command specified as parameter.\n\tExample:\n\thelp echo"
}

