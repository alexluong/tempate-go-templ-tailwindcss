package web

import (
	"embed"
	"net/http"
)

//go:embed dist
var distDir embed.FS

// DistDirFS contains the embedded dist directory files
var DistDirFS = http.FS(distDir)
