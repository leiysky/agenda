package commands

import (
	"github.com/leiysky/agenda/services/store"
	"github.com/spf13/cobra"
)

type cmOptions struct {
	refs []string
}

func newCreateMeetingCommand() *cobra.Command {
	var options cmOptions
	cmd := &cobra.Command{
		Use:   "cm MEETING [...ARGUMENTS]",
		Short: "create a meeting with arguments",
		Args: func(cmd *cobra.Command, args []string) error {
			// if len(args) < 3 {
			// 	return errors.New("requires at least 3 arg")
			// }
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
	meeting := &store.Meeting{}
	meeting.Title = opts.refs[0]
	return store.CreateMeeting(meeting)
}
