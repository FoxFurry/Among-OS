package command

type echo struct {}

func (e *echo) Execute(parameters []string) string{
	if len(parameters) == 0 {
		return "Not enough arguments. Expected at least 1"
	}
	return parameters[0]
}

func (e *echo) Help() string {
	return "Echo is a command tool used for displaying lines of text or string which are passed as arguments on the command line"
}
