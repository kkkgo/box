package outbound

import (
	"context"

	"github.com/sagernet/sing-box/adapter"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
)

func New(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.Outbound) (adapter.Outbound, error) {
	var metadata *adapter.InboundContext
	if tag != "" {
		ctx, metadata = adapter.AppendContext(ctx)
		metadata.Outbound = tag
	}
	if options.Type == "" {
		return nil, E.New("missing outbound type")
	}
	ctx = ContextWithTag(ctx, tag)
	switch options.Type {
	case C.TypeDirect:
		return NewDirect(router, logger, tag, options.DirectOptions)
	case C.TypeBlock:
		return NewBlock(logger, tag), nil
	case C.TypeDNS:
		return NewDNS(router, tag), nil
	case C.TypeSOCKS:
		return NewSocks(router, logger, tag, options.SocksOptions)
	case C.TypeHTTP:
		return NewHTTP(ctx, router, logger, tag, options.HTTPOptions)
	case C.TypeSelector:
		return NewSelector(ctx, router, logger, tag, options.SelectorOptions)
	case C.TypeURLTest:
		return NewURLTest(ctx, router, logger, tag, options.URLTestOptions)
	default:
		return nil, E.New("unknown outbound type: ", options.Type)
	}
}
