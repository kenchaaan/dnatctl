package list

import (
	"fmt"
	"github.com/kenchaaan/dnatctl/pkg/util"
	"github.com/spf13/cobra"
)

func NewListCommand(stream util.IOStream) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list shrt",
		Long:  "list long",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("ilst")
		},
	}

	return cmd
}
