package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/leiysky/agenda/services/store"
	"github.com/spf13/cobra"
)

type rmOptions struct {
	refs []string
	all  bool
}

func newQueryMeetingCommand() *cobra.Command {
	var options rmOptions
	cmd := &cobra.Command{
		Use:   "rm START_TIME END_TIME",
		Short: "query meetings",
		RunE: func(cmd *cobra.Command, args []string) error {
			options.refs = args
			return runQueryMeetingCommand(options)
		},
	}
	cmd.Flags().BoolVarP(&options.all, "all", "A", false, "retrieve all meetings")
	return cmd
}

func runQueryMeetingCommand(opts rmOptions) error {
	if !store.IsLoggedIn() {
		return errors.New("not authenticated")
	}
	if !opts.all && len(opts.refs) < 2 {
		return errors.New("requires at least 2 arguments")
	}
	meetings, err := store.GetAllMeetings()
	var filter store.MeetingsType
	if err != nil {
		return err
	}
	if !opts.all {
		start, _ := store.NewDate(opts.refs[0])
		end, _ := store.NewDate(opts.refs[1])
		for _, one := range meetings {
			stime := one.StartTime
			etime := one.EndTime
			if end.Between(stime, etime) || start.Between(stime, etime) || (start.Gt(stime) && end.Lt(etime)) {
				filter = append(filter, one)
			}
		}
	} else {
		filter = meetings
	}
	os.Stdout.Write([]byte("TITLE               START_TIME          END_TIME            PARTICIPATORS\n"))
	for _, one := range filter {
		if _, err := fmt.Printf("%-20s%-20s%-20s", one.Title, store.DateToString(one.StartTime), store.DateToString(one.EndTime)); err != nil {
			return err
		}
		for _, p := range one.Participators {
			fmt.Print(p)
		}
		fmt.Println()
	}
	return nil
}
