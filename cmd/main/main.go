package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/alexluong/template-go-templ-tailwindcss/internal/server"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(
	ctx context.Context,
	args []string,
	getenv func(string) string,
) error {
	return server.Run(ctx, &server.ServerConfig{
		Port:      8090,
		DistEmbed: !isDevelopment(args, getenv),
	})
}

func isDevelopment(args []string, getenv func(string) string) bool {
	if strings.Contains(args[0], "tmp/") {
		return true
	}
	if getenv("APP_ENV") == "development" {
		return true
	}
	return false
}
