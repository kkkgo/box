package constant

const (
	TypeRedirect = "redirect"
	TypeTProxy   = "tproxy"
	TypeDirect   = "direct"
	TypeBlock    = "block"
	TypeDNS      = "dns"
	TypeSOCKS    = "socks"
	TypeHTTP     = "http"
	TypeMixed    = "mixed"
)

const (
	TypeSelector = "selector"
	TypeURLTest  = "urltest"
)

func ProxyDisplayName(proxyType string) string {
	switch proxyType {
	case TypeDirect:
		return "Direct"
	case TypeBlock:
		return "Block"
	case TypeDNS:
		return "DNS"
	case TypeSOCKS:
		return "SOCKS"
	case TypeHTTP:
		return "HTTP"
	case TypeSelector:
		return "Selector"
	case TypeURLTest:
		return "URLTest"
	default:
		return "Unknown"
	}
}
