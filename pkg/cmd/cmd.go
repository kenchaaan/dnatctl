package cmd

import (
	"github.com/spf13/cobra"
	"github.com/kenchaaan/dnatctl/pkg/cmd/version"
)

var (
	src string
)

// NewDnatctlCommand creates the `dnatctl` commnad and its nested children.
func NewDnatctlCommand() *cobra.Command {
	// Parent command to which all subcommansd are added.
	cmds := &cobra.Command{
		Use:   "dnatctl",
		Short: "dnatctl controls the DNAT configurations of NSX-T.",
		Long: "dnatctl controls the DNAT configurations of NSX-T",
		Run: RunHelp,
	}


	cmds.AddCommand(version.NewVersionCmmand())

	return cmds
}

func RunHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}