package cmd

import (
	"github.com/kenchaaan/dnatctl/pkg/cmd/version"
	"github.com/kenchaaan/dnatctl/pkg/util"
	"github.com/kenchaaan/dnatctl/pkg/cmd/delete"
	"github.com/kenchaaan/dnatctl/pkg/cmd/create"
	"github.com/kenchaaan/dnatctl/pkg/cmd/list"
	"github.com/spf13/cobra"
	"os"
)

type Steram struct {

}

// NewDfaultDnatctlCommand creates the `dnatctl` command with default arguments.
func NewDeafultDnatctlCommand() *cobra.Command {
	return NewDnatctlCommand(util.IOStream{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
}

// NewDnatctlCommand creates the `dnatctl` commnad and its nested children.
func NewDnatctlCommand(stream util.IOStream) *cobra.Command {
	// Parent command to which all subcommands are added.
	cmds := &cobra.Command{
		Use:   "dnatctl",
		Short: "dnatctl controls the DNAT configurations of NSX-T.",
		Long:  "dnatctl controls the DNAT configurations of NSX-T",
		Run:   RunHelp,
	}

	cmds.AddCommand(version.NewVersionCmmand(stream))
	cmds.AddCommand(list.NewListCommand(stream))
	cmds.AddCommand(create.NewCreateCommand(stream))
	cmds.AddCommand(delete.NewDeleteCommand(stream))


	return cmds
}

func RunHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
