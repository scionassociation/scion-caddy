package scion

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/scionproto/scion/pkg/snet"
)

func init() {
	caddy.RegisterModule(Middleware{})
}

type Middleware struct{}

// CaddyModule returns the Caddy module information.
func (Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.scion",
		New: func() caddy.Module { return new(Middleware) },
	}
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	if _, err := snet.ParseUDPAddr(r.RemoteAddr); err == nil {
		r.Header.Add("X-SCION", "on")
		r.Header.Add("X-SCION-Remote-Addr", r.RemoteAddr)
	} else {
		r.Header.Add("X-SCION", "off")
	}
	return next.ServeHTTP(w, r)
}

var _ caddyhttp.MiddlewareHandler = (*Middleware)(nil)
