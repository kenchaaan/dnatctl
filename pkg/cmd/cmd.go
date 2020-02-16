package cmd

import (
	"fmt"
	"github.com/kenchaaan/dnatctl/pkg/cmd/create"
	"github.com/kenchaaan/dnatctl/pkg/cmd/delete"
	"github.com/kenchaaan/dnatctl/pkg/cmd/list"
	"github.com/kenchaaan/dnatctl/pkg/cmd/version"
	"github.com/kenchaaan/dnatctl/pkg/dnatclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)


// NewDfaultDnatctlCommand creates the `dnatctl` command with default arguments.
func NewDeafultDnatctlCommand() *cobra.Command {
	return NewDnatctlCommand(dnatclient.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
}

// NewDnatctlCommand creates the `dnatctl` commnad and its nested children.
func NewDnatctlCommand(stream dnatclient.IOStreams) *cobra.Command {
	cobra.OnInitialize(initConfig)

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

func initConfig() {
	home, _ := os.Getwd()
	viper.AddConfigPath(home)
	viper.SetConfigName(".dnatctl")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		fmt.Println(err)
	}

	s := viper.GetStringMapString("nsxt")
	m := viper.GetStringMapString("pair_globals_to_pseudos")
	dnatclient.Initialize(s["endpoint"], s["username"], s["password"], s["logical_router_id"], m)
}