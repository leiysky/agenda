package commands

import (
	"errors"

	"github.com/leiysky/agenda/services/store"
	"github.com/spf13/cobra"
)

type emOptions struct {
	refs []string
}

func newExitMeetingCommand() *cobra.Command {
	var options emOptions
	cmd := &cobra.Command{
		Use:   "em TITLE",
		Short: "exit meeting",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires at least 1 arguments")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			options.refs = args
			return runExitMeetingCommand(options)
		},
	}

	return cmd
}

func runExitMeetingCommand(opts emOptions) error {
	if !store.IsLoggedIn() {
		return errors.New("not authenticated")
	}
	client, err := store.GetClient()
	if err != nil {
		return err
	}
	meetings := &client.DB.Collection.Meetings
	var meeting *store.MeetingType
	for _, one := range *meetings {
		if one.Title == opts.refs[0] {
			meeting = &one
		}
	}
	if meeting == nil {
		return errors.New("meeting doesn't exist")
	}
	user := store.GetCurrentUser()
	if user.CanTakePartIn(meeting) {
		if err := meeting.AddParticipator(user.Username); err != nil {
			return err
		}
	}
	return nil
}
