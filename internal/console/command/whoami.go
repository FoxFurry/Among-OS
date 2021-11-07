package command

type whoami struct {}

func (e *whoami) Execute(parameters []string) string{
	if len(parameters) != 0 {
		return "Not expected any parameters for <$ whoami>"
	}
	return "Amogus"
}

func (e *whoami) Help() string {
	return "whoami returns current user's name. Since there is no user system - it returns always Amogus"
}

