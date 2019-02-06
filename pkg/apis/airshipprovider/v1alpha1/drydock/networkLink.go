// Code generated by schema-generate. DO NOT EDIT.

package drydock

// Bonding
type Bonding struct {
	DownDelay float64 `json:"down_delay,omitempty"`
	Hash      string  `json:"hash,omitempty"`
	Mode      string  `json:"mode,omitempty"`
	MonRate   float64 `json:"mon_rate,omitempty"`
	PeerRate  string  `json:"peer_rate,omitempty"`
	UpDelay   float64 `json:"up_delay,omitempty"`
}

// Labels
type Labels2 struct {
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}

// Root
type NetworkLinkSpec struct {
	AllowedNetworks []string  `json:"allowed_networks,omitempty"`
	Bonding         *Bonding  `json:"bonding,omitempty"`
	Labels          *Labels2  `json:"labels,omitempty"`
	Linkspeed       string    `json:"linkspeed,omitempty"`
	Mtu             float64   `json:"mtu,omitempty"`
	Trunking        *Trunking `json:"trunking,omitempty"`
}

// Trunking
type Trunking struct {
	DefaultNetwork string `json:"default_network,omitempty"`
	Mode           string `json:"mode,omitempty"`
}