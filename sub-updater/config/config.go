package config

const (
	CurrentVer string = "v0.1.0"
	ConfigName string = "sub-updater.yaml"
)

// DO NOTE: THE sub-sub-updater.yaml should and only should put as the same folder of binary.

var ClientConf *ClientConfig
var OriISPClashConf *ClashConfig

type ClientConfig struct {
	NodeProvider      map[string]string   `yaml:"providers"`
	UseProvider       string              `yaml:"use-provider"`
	MmdbDwnldURL      string              `yaml:"mmdb-download-url"`
	CoreDwnldURL      string              `yaml:"core-download-url"`
	DashboardDwnldURL string              `yaml:"dashboard-download-url"`
	OriginalClashConf *ClientOriClashConf `yaml:"clashori"`
	Rules2Insert 	[]string  `yaml:"rules-insert,omitempty"`
	ClashCorePath     string              `yaml:"clash-core-path"`
	ClashConfPath     string              `yaml:"clash-config-path"`
	CaptivePortal 	  string 	`yaml:"connectivity-portal"`
}

type ClientOriClashConf struct {
	Inbound      *ClashInbound
	Controller   *ClashController
	DNS          *ClashDNS `yaml:"dns"`
}
