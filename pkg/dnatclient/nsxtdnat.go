package dnatclient

import (
	nsxt "github.com/vmware/go-vmware-nsxt"
	"github.com/vmware/go-vmware-nsxt/common"
	"github.com/vmware/go-vmware-nsxt/manager"
)

type DnatClientConfiguration struct {
	LogicalRouterId string
	IpMaps map[string]string
	nsxClient *nsxt.APIClient
}

const (
	scope = "wpjp/exposure"
) 

var d = new(DnatClientConfiguration)

// TODO(kenji-kondo): Define validation functions:
//   * Force to map like key:value -> globalIP:pseudoGlobalIP.
//   * Force all params are fulfilled.
func Initialize(nsxthost string, username string, password string, logicalRouterId string, ipPairs map[string]string) error {
	return d.Initialize(nsxthost, username, password, logicalRouterId, ipPairs)
}
func ( d *DnatClientConfiguration) Initialize(nsxthost string, username string, password string, logicalRouterId string, ipPairs map[string]string) error {
	cfg := &nsxt.Configuration{
		BasePath:  "/api/v1",
		Host:      nsxthost,
		Scheme:    "https",
		UserAgent: "dnatctl/1.0.0/go",
		UserName:  username,
		Password:  password,
		Insecure:  true,
		RemoteAuth: false,
		DefaultHeader: make(map[string]string),
		RetriesConfiguration: nsxt.ClientRetriesConfiguration{
			RetryMinDelay: 1000,
		},
	}
	d.nsxClient, _ = nsxt.NewAPIClient(cfg)
	d.LogicalRouterId = logicalRouterId
	d.IpMaps = ipPairs
	return nil
}

// TODO(kenji-kondo): Enable it to list with the scope/tag and return as a more
//   user friendly result format.
func ListDnatConfigurations() (resrult []manager.NatRule) {
	rr, _, _ := d.nsxClient.LogicalRoutingAndServicesApi.ListNatRules(d.nsxClient.Context, d.LogicalRouterId, nil)
	result := rr.Results
	var r []manager.NatRule

	for _, i := range result {
		if len(i.Tags) != 0 && i.Tags[0].Scope == scope {
			r = append(r, i)
		}
	}
	return r
}

func AddDnatConfiguration(displayName string, globalIp string, translatedIp string) error {
	pseudoIp := d.IpMaps[globalIp]
	var f = func(ip string) manager.NatRule{
		return manager.NatRule{
			DisplayName:             displayName,
			Action:                  "DNAT",
			Enabled:                 true,
			MatchDestinationNetwork: ip,
			MatchService:            nil,
			RulePriority:            100,
			TranslatedNetwork:       translatedIp,
			FirewallMatch:           "MATCH_EXTERNAL_ADDRESS",
			Tags: []common.Tag{{
				Scope: scope,
				Tag:   translatedIp,
			}},
		}
	}
	_, _, err := d.nsxClient.LogicalRoutingAndServicesApi.AddNatRule(d.nsxClient.Context, d.LogicalRouterId, f(globalIp))
	_, _, err = d.nsxClient.LogicalRoutingAndServicesApi.AddNatRule(d.nsxClient.Context, d.LogicalRouterId, f(pseudoIp))
	return err
}

func DeleteDnatConfiguration(ip string) error {
	r := ListDnatConfigurations()
	var ids []string
	for _, i := range r {
		if len(i.Tags) != 0 && i.Tags[0].Tag == ip {
			ids = append(ids, i.Id)
		}
	}
	for _, i := range ids {
		d.nsxClient.LogicalRoutingAndServicesApi.DeleteNatRule(d.nsxClient.Context, d.LogicalRouterId, i)
	}
	return nil
}