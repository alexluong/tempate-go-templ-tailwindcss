package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

type ServerConfig struct {
	Host string
	Port string
}

func Run(ctx context.Context, config *ServerConfig) error {
	srv := newServer()
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}
	log.Printf("listening on %s\n", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		return err
	}
	return nil
}

func newServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	return mux
}
