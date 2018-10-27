package commands

import (
	"errors"

	"github.com/leiysky/agenda/services/store"

	"github.com/spf13/cobra"
)

type cpOptions struct {
	refs []string
}

func newCreateParticipatorCommand() *cobra.Command {
	var options cpOptions
	cmd := &cobra.Command{
		Use:   "cp MEETING [...PARTICIPATORS]",
		Short: "",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("requires at least 2 arguments")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			options.refs = args
			return runCreateParticipatorCommand(options)
		},
	}

	return cmd
}

func runCreateParticipatorCommand(opts cpOptions) error {
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
	participators := opts.refs[0:]
	for _, one := range participators {
		user := store.UserType{one, ""}
		if user.CanTakePartIn(meeting) {
			if err := meeting.AddParticipator(one); err != nil {
				return err
			}
		}
	}
	if err := client.Commit(); err != nil {
		return err
	}
	client.Dump()
	return nil
}
