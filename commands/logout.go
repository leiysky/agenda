package commands

import (
	"github.com/leiysky/agenda/services/store"

	"github.com/spf13/cobra"
)

type logoutOptions struct {
	refs []string
}

func newLogoutCommand() *cobra.Command {
	var options logoutOptions
	cmd := &cobra.Command{
		Use:   "login USERNAME PASSWORD",
		Short: "login",
		RunE: func(cmd *cobra.Command, args []string) error {
			options.refs = args
			return runLogoutCommand(options)
		},
	}

	return cmd
}

func runLogoutCommand(opts logoutOptions) error {
	user := store.UserType{"", ""}
	return store.UpdateLoginState(user)
}
