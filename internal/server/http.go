package server

import (
	cardv1 "foxcard/api/card/v1"
	"foxcard/internal/conf"
	"foxcard/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server,
	ac *conf.Auth,
	card *service.CardService,
	logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			MiddlewareAuth(ac),
			MiddlewareSign(),
		),
		http.ResponseEncoder(ResponseEncoder),
		http.ErrorEncoder(ErrorEncoder),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	cardv1.RegisterCardHTTPServer(srv, card)

	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)

	return srv
}
