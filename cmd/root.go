package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ksbuilder",
		Short: "ksbuilder is a command line interface for KubeSphere extension system",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(cmd.UsageString())

			return nil
		},
	}

	cmd.AddCommand(newVersionCmd(version))  // version subcommand
	cmd.AddCommand(newProjectCmd())         // init_project subcommand
	cmd.AddCommand(newExtensionCmd())       // create_extension subcommand
	cmd.AddCommand(installExtensionCmd())   // publish_extension subcommand
	cmd.AddCommand(uninstallExtensionCmd()) // uninstall_extension subcommand
	cmd.AddCommand(updateExtensionCmd())    // update_extension subcommand

	return cmd
}

// Execute invokes the command.
func Execute(version string) error {
	if err := NewRootCmd(version).Execute(); err != nil {
		return fmt.Errorf("error executing root command: %w", err)
	}

	return nil
}
