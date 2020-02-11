package list

import (
	"encoding/json"
	"fmt"
	"github.com/kenchaaan/dnatctl/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ListOptions struct {
	Verbose bool

	IOStream util.IOStreams
}

func NewListOptions(stream util.IOStreams) *ListOptions{
	return &ListOptions{
		Verbose:  false,
		IOStream: stream,
	}
}

func NewListCommand(stream util.IOStreams) *cobra.Command {
	o := NewListOptions(stream)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all DNAT configurations you did",
		Long:  "List all DNAT configurations you did",
		Run: func(cmd *cobra.Command, args []string) {
			o.Run(cmd, args)
		},
	}

	cmd.Flags().BoolVarP(&o.Verbose, "verbose", "v", o.Verbose, "verbose output")

	return cmd
}

func (o *ListOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

func (o *ListOptions) Run(cmd *cobra.Command, args []string) error {
	//fmt.Println("listed", o.Verbose)
	//r := viper.Get("mappingGlobalToPseudo")
	//a, _ := json.Marshal(r)
	//fmt.Println(string(a))
	//return fmt.Errorf("jfjf")
	v := viper.GetStringMapString("nsxt")
	fmt.Println(v["logicalRouterId"])
	d := util.DnatConfigurations{
		LogicalRouterId: v["logicalrouterid"],
	}
	results := util.ListNsxtDnatConfigurations(d)
	r, _ := json.Marshal(results)
	fmt.Println(string(r))
	fmt.Println(results)
	return nil
}

