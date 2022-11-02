package main

import (
	"net/http"

	"github.com/wtrb/fxhello/handler"
	"github.com/wtrb/fxhello/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			zap.NewExample,
			AsRoute(handler.NewEchoHandler),
			AsRoute(handler.NewHelloHandler),
			fx.Annotate(
				NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			server.New,
		),
		fx.Invoke(
			func(*http.Server) {},
		),
	).Run()
}

type Route interface {
	http.Handler
	Pattern() string
}

func NewServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
