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

	Hostname string

	IOStream dnatclient.IOStreams
}

func NewDeleteOptions(streams dnatclient.IOStreams) *DeleteOptions {
	return &DeleteOptions{
		Verbose:   false,
		CanDelete: false,
		Hostname:  "",
		IOStream:  streams,
	}
}

func NewDeleteCommand(streams dnatclient.IOStreams) *cobra.Command {
	o := NewDeleteOptions(streams)

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete the DNAT configuration identified by hostname",
		Long:  "Delete the DNAT configuration identified by hostname",
		Run: func(cmd *cobra.Command, args []string) {
			o.Validate(cmd, args)
			o.Run(cmd, args)
		},
	}

	cmd.Flags().BoolVar(&o.CanDelete, "non-interactive", o.CanDelete, "do deletion with non interactive")
	cmd.Flags().StringVarP(&o.Hostname, "external-hostname", "n", o.Hostname, "(REQUIRED) IP of the target DNAT configuration")
	cmd.MarkFlagRequired("external-hostname")


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
	dnatclient.DeleteDnatConfiguration(o.Hostname)
	return nil
}