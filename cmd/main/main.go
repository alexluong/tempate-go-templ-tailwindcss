package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alexluong/template-go-templ-tailwindcss/internal/server"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	return server.Run(ctx, &server.ServerConfig{
		Host: "localhost",
		Port: "8090",
	})
}
