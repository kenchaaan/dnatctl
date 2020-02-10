package version

import (
	"encoding/json"
	"fmt"
	"github.com/kenchaaan/dnatctl/pkg/util"

	"github.com/spf13/cobra"
)

var (
	GitVersion = "1.0"
	GitCommit  = "aaaaaa"
	BuildDate  = "1970-01-01T00:00:00Z"
	GoVersion  = "1.5"
)

type Info struct {
	GitVersion string `json:"version"`
	GitCommit  string `json:"commitId"`
	BuildDate  string `json:"BuildDate"`
	GoVersion  string `json:"GoVersion"`
}

func NewVersionCmmand(stream util.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of dnatctl",
		Long:  "Print the version number of dnatctl",
		Run: func(cmd *cobra.Command, args []string) {
			r, err := json.Marshal(Get())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(r))
		},
	}

	return cmd
}

func Get() Info {
	return Info{
		GitVersion: GitVersion,
		GitCommit:  GitCommit,
		BuildDate:  BuildDate,
		GoVersion:  GoVersion,
	}
}
