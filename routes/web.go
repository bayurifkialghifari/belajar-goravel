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
			"server_time": time.Now().Format("2006-01-02 15:04:05 MST -0700"), // includes zone name and offset
			"version":     support.Version,
		})
	}).Name("health.check")
}
