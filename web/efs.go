package web

import (
	"embed"
)

//go:embed "static" "templates"
var embeddedFiles embed.FS

func GetEmbeddedFiles() embed.FS {
	return embeddedFiles
}
