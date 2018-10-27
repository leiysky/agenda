package commands

import (
	"errors"

	"github.com/leiysky/agenda/services/store"
	"github.com/spf13/cobra"
)

type cmOptions struct {
	refs []string
}

func newCreateMeetingCommand() *cobra.Command {
	var options cmOptions
	cmd := &cobra.Command{
		Use:   "cm TITLE START_TIME END_TIME [...PARTICIPATORS]",
		Short: "create a meeting with arguments",
		Long: `Date format is similar to UNIX time
		Format: YYYY:MM:DDThh:mm
		Example: 2018-10-24T13:59 `,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 3 {
				return errors.New("requires at least 3 arguments")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			options.refs = args
			return runCreateMeetingsCommand(options)
		},
	}

	return cmd
}

func runCreateMeetingsCommand(opts cmOptions) error {
	if !store.IsLoggedIn() {
		return errors.New("not authenticated")
	}
	meeting := store.MeetingType{}
	meeting.Title = opts.refs[0]
	start, err := store.NewDate(opts.refs[1])
	if err != nil {
		return err
	}
	meeting.StartTime = *start
	end, err := store.NewDate(opts.refs[2])
	if err != nil {
		return err
	}
	meeting.EndTime = *end
	meeting.Participators = opts.refs[3:]
	meeting.Participators = append(meeting.Participators, store.GetCurrentUser().Username)
	for _, one := range meeting.Participators {
		user := store.UserType{one, ""}
		if !user.IsExist() {
			return errors.New(one + " is a invalid user")
		}
		if !user.CanTakePartIn(&meeting) {
			return errors.New(one + " can't take part in the meeting")
		}
	}
	return store.CreateMeeting(meeting)
}
