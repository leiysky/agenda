package commands

import (
	"errors"
	"regexp"

	"github.com/leiysky/agenda/services/store"

	"github.com/spf13/cobra"
)

type registerOptions struct {
	refs []string
}

func newRegisterCommand() *cobra.Command {
	options := registerOptions{}
	cmd := &cobra.Command{
		Use:   "register USERNAME PASSWORD",
		Short: "register an account",
		Long:  "username should be started with _ or alphabet, and keep its length between 6 and 12",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("requires at least 2 arguments")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			options.refs = args
			return runRegisterCommand(options)
		},
	}

	return cmd
}

func runRegisterCommand(opts registerOptions) error {
	user := store.UserType{}
	user.Username = opts.refs[0]
	if !validateUsername(user.Username) {
		return errors.New("invalid username")
	}
	user.Password = opts.refs[1]
	return store.CreateUser(user)
}

func validateUsername(username string) bool {
	if istrue, _ := regexp.Match("[_a-zA-Z][_a-zA-Z0-9]{5,11}", []byte(username)); istrue == false {
		return false
	}
	return true
}
