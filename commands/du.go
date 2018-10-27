package commands

import (
	"errors"

	"github.com/leiysky/agenda/services/store"
	"github.com/spf13/cobra"
)

type duOptions struct {
	refs []string
}

func newDeleteUserCommand() *cobra.Command {
	var options duOptions
	cmd := &cobra.Command{
		Use:   "du USERNAME",
		Short: "delete users",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires at least 1 arguments")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			options.refs = args
			return runDeleteUserCommand(options)
		},
	}

	return cmd
}

func runDeleteUserCommand(opts duOptions) error {
	if !store.IsLoggedIn() {
		return errors.New("not authenticated")
	}
	username := opts.refs[0]
	return store.DeleteUserByName(username)
}
