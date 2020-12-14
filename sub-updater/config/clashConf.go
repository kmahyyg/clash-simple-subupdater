package config

// Copied from github.com/dreamacro/clash/config, simplified version
// Removed some configuration and simplified some type to base type

// ClashConfig read in the original config of clash and parse
type ClashConfig struct {
	General    *ClashGeneral `yaml:"-"`
	DNS        *ClashDNS `yaml:"-"`
	Controller *ClashController `yaml:"-"`
	Inbound    *ClashInbound	`yaml:"-"`
	NodeNRoute *ClashNodeAndRoute `yaml:"-"`
}

type ClashGeneral struct {
	Mode     string `yaml:"mode"`
	LogLevel string `yaml:"log-level"`
	IPv6     bool   `yaml:"ipv6,omitempty"`
}

type ClashController struct {
	ExternalController string `yaml:"external-controller,omitempty"`
	ExternalUI         string `yaml:"external-ui,omitempty"`
	Secret             string `yaml:"secret,omitempty"`
}

type ClashInbound struct {
	Port        int    `yaml:"port,omitempty"`
	SocksPort   int    `yaml:"socks-port,omitempty"`
	RedirPort   int    `yaml:"redir-port"`
	TProxyPort  int    `yaml:"tproxy-port,omitempty"`
	MixedPort   int    `yaml:"mixed-port"`
	AllowLan    bool   `yaml:"allow-lan"`
	BindAddress string `yaml:"bind-address,omitempty"`
}

type ClashDNS struct {
	Enable            bool                 `yaml:"enable"`
	IPv6              bool                 `yaml:"ipv6"`
	NameServer        []string             `yaml:"nameserver"`
	Fallback          []string             `yaml:"fallback,omitempty"`
	FallbackFilter    *ClashFallbackFilter `yaml:"fallback-filter,omitempty"`
	Listen            string               `yaml:"listen"`
	EnhancedMode      string               `yaml:"enhanced-mode"`
	DefaultNameserver []string             `yaml:"default-nameserver"`
	FakeIPRange       string               `yaml:"fake-ip-range"`
	FakeIPFilter      []string             `yaml:"fake-ip-filter,omitempty"`
	UseHosts          bool                 `yaml:"use-hosts,omitempty"`
	Hosts             map[string]string    `yaml:"hosts,omitempty"`
}

type ClashFallbackFilter struct {
	GeoIP  bool     `yaml:"geoip"`
	IPCIDR []string `yaml:"ipcidr"`
	Domain []string `yaml:"domain"`
}

type ClashNodeAndRoute struct {
	ProxyProvider map[string]map[string]interface{} `yaml:"proxy-providers,omitempty"`
	Proxy         []map[string]interface{}          `yaml:"proxies"`
	ProxyGroup    []map[string]interface{}          `yaml:"proxy-groups,omitempty"`
	Rule          []string                          `yaml:"rules"`
}
