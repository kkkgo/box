package option

import (
	C "github.com/sagernet/sing-box/constant"
	E "github.com/sagernet/sing/common/exceptions"
	"github.com/sagernet/sing/common/json"
)

type _Inbound struct {
	Type            string                  `json:"type"`
	Tag             string                  `json:"tag,omitempty"`
	RedirectOptions RedirectInboundOptions  `json:"-"`
	TProxyOptions   TProxyInboundOptions    `json:"-"`
	DirectOptions   DirectInboundOptions    `json:"-"`
	SocksOptions    SocksInboundOptions     `json:"-"`
	HTTPOptions     HTTPMixedInboundOptions `json:"-"`
	MixedOptions    HTTPMixedInboundOptions `json:"-"`
}

type Inbound _Inbound

func (h Inbound) MarshalJSON() ([]byte, error) {
	var v any
	switch h.Type {
	case C.TypeRedirect:
		v = h.RedirectOptions
	case C.TypeTProxy:
		v = h.TProxyOptions
	case C.TypeDirect:
		v = h.DirectOptions
	case C.TypeSOCKS:
		v = h.SocksOptions
	case C.TypeHTTP:
		v = h.HTTPOptions
	case C.TypeMixed:
		v = h.MixedOptions
	default:
		return nil, E.New("unknown inbound type: ", h.Type)
	}
	return MarshallObjects((_Inbound)(h), v)
}

func (h *Inbound) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, (*_Inbound)(h))
	if err != nil {
		return err
	}
	var v any
	switch h.Type {
	case C.TypeRedirect:
		v = &h.RedirectOptions
	case C.TypeTProxy:
		v = &h.TProxyOptions
	case C.TypeDirect:
		v = &h.DirectOptions
	case C.TypeSOCKS:
		v = &h.SocksOptions
	case C.TypeHTTP:
		v = &h.HTTPOptions
	case C.TypeMixed:
		v = &h.MixedOptions
	default:
		return E.New("unknown inbound type: ", h.Type)
	}
	err = UnmarshallExcluded(bytes, (*_Inbound)(h), v)
	if err != nil {
		return E.Cause(err, "inbound options")
	}
	return nil
}

type InboundOptions struct {
	SniffEnabled              bool           `json:"sniff,omitempty"`
	SniffOverrideDestination  bool           `json:"sniff_override_destination,omitempty"`
	SniffTimeout              Duration       `json:"sniff_timeout,omitempty"`
	DomainStrategy            DomainStrategy `json:"domain_strategy,omitempty"`
	UDPDisableDomainUnmapping bool           `json:"udp_disable_domain_unmapping,omitempty"`
}

type ListenOptions struct {
	Listen                      *ListenAddress `json:"listen,omitempty"`
	ListenPort                  uint16         `json:"listen_port,omitempty"`
	TCPFastOpen                 bool           `json:"tcp_fast_open,omitempty"`
	TCPMultiPath                bool           `json:"tcp_multi_path,omitempty"`
	UDPFragment                 *bool          `json:"udp_fragment,omitempty"`
	UDPFragmentDefault          bool           `json:"-"`
	UDPTimeout                  int64          `json:"udp_timeout,omitempty"`
	ProxyProtocol               bool           `json:"proxy_protocol,omitempty"`
	ProxyProtocolAcceptNoHeader bool           `json:"proxy_protocol_accept_no_header,omitempty"`
	Detour                      string         `json:"detour,omitempty"`
	InboundOptions
}