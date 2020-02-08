package create

import (
	"fmt"
	"github.com/kenchaaan/dnatctl/pkg/util"
	"github.com/spf13/cobra"
)

func NewCreateCommand(stream util.IOStream) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create shrt",
		Long:  "create long",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("create")
		},
	}

	return cmd
}
