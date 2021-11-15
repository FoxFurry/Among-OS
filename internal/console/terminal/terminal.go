package terminal

import (
	"bufio"
	"context"
	"fmt"
	"github.com/foxfurry/university/sommip/internal/console/command"
	"io"
	"os"
	"strings"
)

type ITerminal interface {
	Start(ctx context.Context)
}

type terminal struct {
	cmdMap      map[string]command.ICommand
	inputPrefix string
}

func NewTerminal() ITerminal {
	return &terminal{
		cmdMap:      command.GetCommands(),
		inputPrefix: ">",
	}
}

func (c *terminal) Start(ctx context.Context) {
	readChannel := c.reader(os.Stdin)

	fmt.Printf(c.inputPrefix + " ")

	for {
		select {
		case cmd := <-readChannel:
			execResult := c.parseInput(cmd)
			fmt.Println(execResult)
			fmt.Printf(c.inputPrefix + " ")
		case <-ctx.Done():
			fmt.Println("Console is closed")
			return
		}
	}
}

func (c *terminal) reader(r io.Reader) <-chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)
		scan := bufio.NewReader(r)
		for {
			data, _ := scan.ReadString('\n')
			lines <- data
		}
	}()
	return lines
}

func (c *terminal) parseInput(input string) string {
	if input == "\n" {
		return "````````````````you have entered nothing give more arguments`````````````````"
	}
	inputFields := strings.Fields(input)

	cmd, ok := c.cmdMap[inputFields[0]]

	if !ok {
		return fmt.Sprintf("command %s not found", inputFields[0])
	}

	result := cmd.Execute(inputFields[1:])

	return result
}
