package command

type hostname struct{}

func (e *hostname) Execute(parameters []string) string {
	if len(parameters) != 0 {
		return "Not expected any parameters for <$ hostname>"
	}
	return "FOX_FURRY_OS-pc"
}

func (e *hostname) Help() string {
	return "hostname show's the system's host name"
}
