package create

import (
	"github.com/kenchaaan/dnatctl/pkg/dnatclient"
	"github.com/spf13/cobra"
)

type CreateOptions struct {

	DisplayName string
	GlobalIp string
	TransportedIp string

	IOStream dnatclient.IOStreams
}

func NewCreateOptions(stream dnatclient.IOStreams) *CreateOptions{
	return &CreateOptions{
		IOStream:     stream,
	}
}

func NewCreateCommand(stream dnatclient.IOStreams) *cobra.Command {
	o := NewCreateOptions(stream)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new DNAT configuration",
		Long:  "Create a new DNAT configuration",
		Run: func(cmd *cobra.Command, args []string) {
			o.Run(cmd, args)
		},
	}

	cmd.Flags().StringVarP(&o.DisplayName, "external-hostname", "n", "", "(REQUIRED) External hostname")
	cmd.Flags().StringVarP(&o.GlobalIp, "global-ip", "g", "", "(REQUIRED) Global IP that accepts via the Internet")
	cmd.Flags().StringVarP(&o.TransportedIp, "transported-ip", "t", "", "(REQUIRED) Transported IP")
	cmd.MarkFlagRequired("external-hostname")
	cmd.MarkFlagRequired("global-ip")
	cmd.MarkFlagRequired("transported-ip")

	return cmd
}

func (o *CreateOptions) Run(cmd *cobra.Command, args []string) error {
	err := dnatclient.AddDnatConfiguration(o.DisplayName, o.GlobalIp, o.TransportedIp)
	return err
}