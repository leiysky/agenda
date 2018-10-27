package commands

import (
	"errors"

	"github.com/leiysky/agenda/services/store"
	"github.com/spf13/cobra"
)

type dmOptions struct {
	refs []string
}

func newDeleteMeetingCommand() *cobra.Command {
	var options dmOptions
	cmd := &cobra.Command{
		Use:   "dm TITLE",
		Short: "delete a meeting",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires at least 1 arguments")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			options.refs = args
			return runDeleteMeetingsCommand(options)
		},
	}

	return cmd
}

func runDeleteMeetingsCommand(opts dmOptions) error {
	if !store.IsLoggedIn() {
		return errors.New("not authenticated")
	}
	meetingName := opts.refs[0]
	return store.DeleteMeetingByName(meetingName)
}
