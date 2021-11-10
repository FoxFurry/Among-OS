package command

import "sync"

type ICommand interface {
	Execute(parameters []string) string
	Help() string
}

var (
	cmdMap     map[string]ICommand
	cmdMapOnce sync.Once
)

func GetCommands() map[string]ICommand {
	cmdMapOnce.Do(func() {
		cmdMap = make(map[string]ICommand)
		cmdMap["echo"] = &echo{}
		cmdMap["help"] = &help{}
		cmdMap["whoami"] = &whoami{}
		cmdMap["uname"] = &uname{}
		cmdMap["hostname"] = &hostname{}
		cmdMap["shutdown"] = &shutdown{}
		cmdMap["gender"] = &nothing{}
	})

	return cmdMap
}
