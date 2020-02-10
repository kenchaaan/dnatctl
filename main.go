package main

import (
	"fmt"
	"github.com/vmware/go-vmware-nsxt"
	"github.com/vmware/go-vmware-nsxt/appdiscovery"
	"net/http"
)



func NewConfiguration() *nsxt.Configuration {
	cfg := &nsxt.Configuration{
		BasePath:  "/api/v1",
		//Host:      "172.28.128.3:8443",
		Host:      "172.28.128.3",
		Scheme:    "https",
		UserAgent: "dnatctl/1.0.0/go",
		UserName:  "admin",
		Password:  "5Wl&Mqkhk2l3",
		Insecure:  true,
		RemoteAuth: false,
		DefaultHeader: map[string]string{"Content-Type": "application/json", "Authorization": "Basic YWRtaW46NVdsJk1xa2hrMmwz"},
	}
	return cfg
}

var (
	dis appdiscovery.AppProfileListResult
	s http.Response
)
func main() {
	cfg := NewConfiguration()
	client, _ := nsxt.NewAPIClient(cfg)
	//_, s, _ := client.AppDiscoveryApi.GetAppProfiles(client.Context)
	logicalRouterId := "c5eba3ce-429c-48ed-a01e-77ece3090731"
	ruleId := "5903"
	_, res, _ := client.LogicalRoutingAndServicesApi.GetNatRule(client.Context, logicalRouterId, ruleId)
	fmt.Println(res)
}