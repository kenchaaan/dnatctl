package delete

import (
	"errors"
	"fmt"
	"github.com/kenchaaan/dnatctl/pkg/dnatclient"
	"github.com/spf13/cobra"
	"github.com/manifoldco/promptui"
)

type DeleteOptions struct {
	Verbose bool
	CanDelete bool

	Ip string

	IOStream dnatclient.IOStreams
}

func NewDeleteOptions(streams dnatclient.IOStreams) *DeleteOptions {
	return &DeleteOptions{
		Verbose:   false,
		CanDelete: false,
		Ip:        "",
		IOStream:  streams,
	}
}

func NewDeleteCommand(streams dnatclient.IOStreams) *cobra.Command {
	o := NewDeleteOptions(streams)

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete short",
		Long:  "Delete long",
		Run: func(cmd *cobra.Command, args []string) {
			o.Validate(cmd, args)
			o.Run(cmd, args)
		},
	}

	cmd.Flags().BoolVar(&o.CanDelete, "non-interactive", o.CanDelete, "do deletion with non interactive")
	cmd.Flags().StringVarP(&o.Ip, "transported-ip", "t", o.Ip, "(REQUIRED) IP of the target DNAT configuration")
	cmd.MarkFlagRequired("transported-ip")


	return cmd
}

func (o *DeleteOptions) Validate(cmd *cobra.Command, args []string) error {
	validate := func(input string) error {
		if input != "Y" && input != "n" {
			return errors.New("Choose Y or n")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:     "Do you want to delete it?",
		Default:   "",
		AllowEdit: false,
		Validate:  validate,
		Mask:      0,
		Templates: nil,
		IsConfirm: false,
		IsVimMode: false,
		Pointer:   nil,
		Stdin:     nil,
		Stdout:    nil,
	}

	if !o.CanDelete {
		resultmap := map[string]bool{"Y": true, "n": false}
		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Error")
		}
		o.CanDelete = resultmap[result]
	}
	return nil
}

func (o *DeleteOptions) Run(cmd *cobra.Command, args []string) error {
	dnatclient.DeleteDnatConfiguration(o.Ip)
	return nil
}