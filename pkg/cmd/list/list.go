package list

import (
	"encoding/json"
	"fmt"
	"github.com/kenchaaan/dnatctl/pkg/dnatclient"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Verbose            bool
	DnatConfigurations dnatclient.DnatClientConfiguration

	IOStream dnatclient.IOStreams
}

func NewListOptions(stream dnatclient.IOStreams) *ListOptions{
	return &ListOptions{
		Verbose:  false,
		IOStream: stream,
	}
}

func NewListCommand(stream dnatclient.IOStreams) *cobra.Command {
	o := NewListOptions(stream)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all DNAT configurations you did",
		Long:  "List all DNAT configurations you did",
		Run: func(cmd *cobra.Command, args []string) {
			o.Run(cmd, args)
		},
	}

	return cmd
}


func (o *ListOptions) Run(cmd *cobra.Command, args []string) error {
	results := dnatclient.ListDnatConfigurations()
	f := func(id string, displayName string, matchDestinationNetwork string, translatedNetwork string) *Results{
		return &Results{
			Id:                      id,
			DisplayName:             displayName,
			MatchDestinationNetwork: matchDestinationNetwork,
			TranslatedNetwork:       translatedNetwork,
		}
	}
	for _, i := range results {
		rrr := f(i.Id, i.DisplayName, i.MatchDestinationNetwork, i.TranslatedNetwork)
		rrrr, _ := json.Marshal(*rrr)
		fmt.Println(string(rrrr))
	}
	return nil
}

type Results struct {
	Id                      string `json:"id"`
	DisplayName             string `json:"display_name"`
	MatchDestinationNetwork string `json:"match_destination_network"`
	TranslatedNetwork       string `json:"translated_network"`
}