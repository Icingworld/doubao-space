package web

import (
	"embed"
	"io/fs"
)

//go:embed dist
var embeddedDist embed.FS

// DistFS is the frontend build output served by the Go application.
var DistFS fs.FS

func init() {
	var err error
	DistFS, err = fs.Sub(embeddedDist, "dist")
	if err != nil {
		panic(err)
	}
}
