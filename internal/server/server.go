package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/alexluong/template-go-templ-tailwindcss/web"
	"github.com/alexluong/template-go-templ-tailwindcss/web/template"
)

type ServerConfig struct {
	Host      string
	Port      string
	DistEmbed bool
}

func Run(ctx context.Context, config *ServerConfig) error {
	srv := newServer(config.DistEmbed)
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

func newServer(distEmbed bool) *http.ServeMux {
	mux := http.NewServeMux()

	if distEmbed {
		distServer := http.FileServer(web.DistDirFS)
		mux.Handle("GET /dist/", distServer)
	} else {
		mux.Handle("GET /dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("web/dist"))))
	}

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		template.HelloWorld().Render(r.Context(), w)
	})

	return mux
}
