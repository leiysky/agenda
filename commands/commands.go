package commands

import (
	"errors"

	"github.com/spf13/cobra"
)

// Create a Agenda CLI
func NewAgendaCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "agenda COMMAND",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires command")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	cmd.AddCommand(
		newCreateMeetingCommand(),
	)
	return cmd
}
