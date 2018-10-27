package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/leiysky/agenda/services/store"
	"github.com/spf13/cobra"
)

type ruOptions struct {
	refs []string
}

func newQueryUserCommand() *cobra.Command {
	var options ruOptions
	cmd := &cobra.Command{
		Use:   "ru",
		Short: "query users",
		RunE: func(cmd *cobra.Command, args []string) error {
			options.refs = args
			return runQueryUserCommand(options)
		},
	}

	return cmd
}

func runQueryUserCommand(opts ruOptions) error {
	if !store.IsLoggedIn() {
		return errors.New("not authenticated")
	}
	users, err := store.GetAllUsers()
	if err != nil {
		return err
	}
	os.Stdout.Write([]byte("USERNAME      \n"))
	for _, one := range users {
		if _, err := fmt.Printf("%s\n", one.Username); err != nil {
			return err
		}
	}
	return nil
}
