package main

import (
	"github.com/leiysky/agenda/commands"
)

func main() {
	cmd := commands.NewAgendaCommand()
	cmd.Execute()
}
