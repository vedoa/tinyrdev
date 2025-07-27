package cli

import "github.com/vedoa/tinyrdev/rpackage/create"

// Execute is the main function reserved to hold all the subcommands if the cli grows
func Execute(v string, args []string) error {
	return create.Execute(v, args)
}
