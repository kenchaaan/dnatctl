package delete

import (
	"fmt"
	"github.com/kenchaaan/dnatctl/pkg/util"
	"github.com/spf13/cobra"
)

func NewDeleteCommand(stream util.IOStream) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete shrt",
		Long:  "Delete long",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("delete")
		},
	}

	return cmd
}