package routes

import (
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
)

func Web() {
	// Health check route
	facades.Route().Get("/health", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("health.tmpl", map[string]any{
			"status":      "OK",
			"server_time": time.Now().Format(time.RFC3339), // includes zone name and offset
			"version":     support.Version,
		})
	}).Name("health.check")
}
