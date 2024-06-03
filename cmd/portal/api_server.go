package main

import (
	"context"
	"net/http"
	"net/http/pprof"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/techx/portal/builder"
	"github.com/techx/portal/client"
	"github.com/techx/portal/config"
	"github.com/techx/portal/middleware"
	"github.com/techx/portal/service"
	"github.com/tylerb/graceful"
	"github.com/urfave/cli/v2"
)

// HTTPAPIServer is a HTTP server for serving APIs.
// This server comes with some default bells & whistles for
// - Logging
// - HTTP Context
// - HTTP Metrics
// - New Relic integration
// - Health Checks
// - Debug/Profile endpoints
// These features are configurable using middlewares.
type HTTPAPIServer struct {
	server *graceful.Server
}

// NewHTTPAPIServer creates a new HTTPAPIServer using a mux.Router.
func NewHTTPAPIServer(cfg config.HTTPAPIConfig, r *mux.Router) *HTTPAPIServer {
	// always attach a health check endpoint
	r.Methods("GET").Path("/ping").HandlerFunc(internalPingHandler)

	// attach pprof & statsviz in debug mode
	if cfg.DebugMode {
		r.HandleFunc("/debug/pprof/", pprof.Index)
		r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		r.HandleFunc("/debug/pprof/profile", pprof.Profile)
		r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		r.HandleFunc("/debug/pprof/trace", pprof.Trace)
		r.HandleFunc("/debug/pprof/goroutine", pprof.Index)
		r.HandleFunc("/debug/pprof/heap", pprof.Index)
		r.HandleFunc("/debug/pprof/threadcreate", pprof.Index)
	}

	// create the http server
	http := http.Server{
		Addr:    cfg.ListenAddr,
		Handler: r,
	}

	// ... and wrap it for graceful shutdowns.
	// graceful shutdowns will ensure abrupt closing/stopping
	// of ongoing requests. Even though, a HAProxy will drain
	// connection. This adds another safety net to ensure critical
	// writes are not missed.
	server := &graceful.Server{
		Timeout: 15 * time.Second,
		Server:  &http,
	}

	return &HTTPAPIServer{
		server: server,
	}
}

// Start starts the API server in a new goroutine.
func (a *HTTPAPIServer) Start() {
	go func(a *HTTPAPIServer) {
		log.Info().Msgf("[API.SERVER] PID %d. Starting server on %s", os.Getpid(), a.server.Addr)

		err := a.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("[API.SERVER] Unhandled server shutdown: %s", err.Error())
		}
	}(a)
}

// Shutdown stops the HTTP server.
func (a *HTTPAPIServer) Shutdown() error {
	return a.server.Shutdown(context.Background())
}

// WaitForShutdown blocks till the server is shutdown.
// Shutdowns can be initiated with a SIGHUP, SIGKILL interupts.
func (a *HTTPAPIServer) WaitForShutdown() bool {
	<-a.server.StopChan()
	log.Info().Msgf("[API.SERVER] Server shutdown complete")

	return true
}

func internalPingHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(("OK")))
}

func startAPIServer(ctx *cli.Context) error {
	log.Info().Msg("starting tech portal server.")

	applicationContext, err := initApplicationContext(ctx)
	if err != nil {
		return err
	}

	// init registry
	clientRegistry := client.NewRegistry(*applicationContext.Config)
	builderRegistry := builder.NewRegistry(*applicationContext.Config, clientRegistry)
	serviceRegistry := service.NewRegistry(*applicationContext.Config, builderRegistry)

	// Create router
	r := NewRouter(applicationContext.Config, serviceRegistry)

	// ... and add middleware
	r.Use(middleware.RecoverMiddleware())

	// ... other middleware go here
	r.Use(middleware.RequestContext())

	// Start API Server
	server := NewHTTPAPIServer(applicationContext.Config.API, r)
	server.Start()

	// Wait for graceful shutdown
	server.WaitForShutdown()

	return nil
}
