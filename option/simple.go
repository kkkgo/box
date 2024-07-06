package option

import "github.com/sagernet/sing/common/auth"

type SocksInboundOptions struct {
	ListenOptions
	Users []auth.User `json:"users,omitempty"`
}

type SocksOutboundOptions struct {
	DialerOptions
	ServerOptions
	Version    string             `json:"version,omitempty"`
	Username   string             `json:"username,omitempty"`
	Password   string             `json:"password,omitempty"`
	Network    NetworkList        `json:"network,omitempty"`
	UDPOverTCP *UDPOverTCPOptions `json:"udp_over_tcp,omitempty"`
}
