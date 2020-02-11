package util

import (
	"fmt"
	nsxt "github.com/vmware/go-vmware-nsxt"
	"github.com/vmware/go-vmware-nsxt/manager"
)

type DnatConfigurations struct {
	GlobalIp string
	PseudoGlobalIp string
	DestinationIp string
	BothId string

	LogicalRouterId string

	nsxClient *nsxt.APIClient
}

var d = new(DnatConfigurations)

func Initialize(cfg *nsxt.Configuration) error { return d.Initialize(cfg)}
func ( d *DnatConfigurations) Initialize(cfg *nsxt.Configuration) error {
	d.nsxClient, _ = nsxt.NewAPIClient(cfg)
	return nil
}

func ListNsxtDnatConfigurations(cfg DnatConfigurations) (result []manager.NatRule) {
	//client, err := nsxt.NewAPIClient(cfg.NsxtCfg)
	//if cfg.BothId != "" {
	//	rr, _, _ := d.nsxClient.LogicalRoutingAndServicesApi.GetNatRule(d.nsxClient.Context, cfg.LogicalRouterId, cfg.BothId)
	//	result = []manager.NatRule{rr}
	//} else {
		rr, _, _ := d.nsxClient.LogicalRoutingAndServicesApi.ListNatRules(d.nsxClient.Context, cfg.LogicalRouterId, nil)
		result = rr.Results
		fmt.Println(result)
	//}
	return
}

