package inbound

import (
	"context"

	"github.com/sagernet/sing-box/adapter"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/experimental/libbox/platform"
	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
)

func New(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.Inbound, platformInterface platform.Interface) (adapter.Inbound, error) {
	if options.Type == "" {
		return nil, E.New("missing inbound type")
	}
	switch options.Type {
	case C.TypeRedirect:
		return NewRedirect(ctx, router, logger, tag, options.RedirectOptions), nil
	case C.TypeTProxy:
		return NewTProxy(ctx, router, logger, tag, options.TProxyOptions), nil
	case C.TypeDirect:
		return NewDirect(ctx, router, logger, tag, options.DirectOptions), nil
	default:
		return nil, E.New("unknown inbound type: ", options.Type)
	}
}
