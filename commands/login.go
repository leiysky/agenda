package commands

import (
	"errors"

	"github.com/leiysky/agenda/services/store"

	"github.com/spf13/cobra"
)

type loginOptions struct {
	refs []string
}

func newLoginCommand() *cobra.Command {
	var options loginOptions
	cmd := &cobra.Command{
		Use:   "login USERNAME PASSWORD",
		Short: "login",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("requires at least 2 arguments")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			options.refs = args
			return runLoginCommand(options)
		},
	}

	return cmd
}

func runLoginCommand(opts loginOptions) error {
	user := store.UserType{opts.refs[0], opts.refs[1]}
	return store.UpdateLoginState(user)
}
