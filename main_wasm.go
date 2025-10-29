//go:build wasm

package main

//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=handler

import (
	"github.com/jpillora/installer/handler"
	"github.com/syumai/workers"
	"github.com/syumai/workers/cloudflare"
)

func main() {
	// Use DefaultConfig directly in WASM environment
	// opts.Parse() doesn't work well in WASM
	c := handler.DefaultConfig

	// Allow environment variable overrides from Workers environment
	// Note: These need to be set in wrangler.jsonc [vars] section
	if defaultUser := cloudflare.Getenv("DEFAULT_USER"); defaultUser != "" && defaultUser != "<undefined>" {
		c.User = defaultUser
	}
	if token := cloudflare.Getenv("GITHUB_TOKEN"); token != "" && token != "<undefined>" {
		c.Token = token
	}
	if forceUser := cloudflare.Getenv("FORCE_USER"); forceUser != "" && forceUser != "<undefined>" {
		c.ForceUser = forceUser
	}
	if forceRepo := cloudflare.Getenv("FORCE_REPO"); forceRepo != "" && forceRepo != "<undefined>" {
		c.ForceRepo = forceRepo
	}
	h := &handler.Handler{Config: c}
	workers.Serve(h)
}
