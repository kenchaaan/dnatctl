module github.com/kenchaaan/dnatctl

go 1.13

require (
	github.com/manifoldco/promptui v0.7.0
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/prometheus/tsdb v0.7.1
	github.com/spf13/afero v1.2.1 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/vmware/go-vmware-nsxt v0.0.0-20200114231430-33a5af043f2e
	golang.org/x/sys v0.0.0-20190502175342-a43fa875dd82 // indirect
	golang.org/x/text v0.3.2 // indirect
)

replace github.com/vmware/go-vmware-nsxt v0.0.0-20200114231430-33a5af043f2e => github.com/vmware/go-vmware-nsxt v0.0.0-847b08bedf5
