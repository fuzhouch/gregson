package cli

import "github.com/spf13/cobra"

type loginOpt struct {
	authType string
	username string
	password string
	cmd      *cobra.Command
}

// NewLoginOption implements a Cobra login subcommand.
func NewLoginOption() *cobra.Command {
	opt := new(loginOpt)
	opt.cmd = &cobra.Command{
		Use:   "login",
		Short: "Client login",
		Long:  "Login by fetching and saving access token",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	opt.cmd.Flags().StringVar(
		&opt.authType,
		"type",
		"ldap",
		"Login type. Supported: ldap (default: ldap)")

	opt.cmd.Flags().StringVar(
		&opt.username,
		"username",
		"",
		"Login username")

	opt.cmd.Flags().StringVar(
		&opt.username,
		"password",
		"",
		"Login password")

	return opt.cmd
}
