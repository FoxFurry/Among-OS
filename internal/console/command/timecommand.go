package command

import "time"

type timecommand struct{}

func (e *timecommand) Execute(parameters []string) string {
	if len(parameters) != 0 {
		return "Not expected any parameters for <$ timecommands>"
	}
	currentTime := time.Now()
	return currentTime.String()
}

func (e *timecommand) Help() string {
	return "it shows realtime on the device in format YYYY-MM-DD HH:MM:SS.SSSSSSS"
}
