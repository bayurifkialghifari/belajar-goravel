package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	contractHttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/http/limit"

	"karuhundeveloper.com/gogo/app/http"
	"karuhundeveloper.com/gogo/routes"
	"karuhundeveloper.com/gogo/routes/api"
)

type RouteServiceProvider struct {
}

func (receiver *RouteServiceProvider) Register(app foundation.Application) {
}

func (receiver *RouteServiceProvider) Boot(app foundation.Application) {
	// Add HTTP middleware
	facades.Route().GlobalMiddleware(http.Kernel{}.Middleware()...)

	receiver.configureRateLimiting()

	// Add routes
	routes.Web()
	api.V1()
}

func (receiver *RouteServiceProvider) configureRateLimiting() {
	// 60 requests per minute
	facades.RateLimiter().For("api", func(ctx contractHttp.Context) contractHttp.Limit {
		return limit.PerMinute(60)
	})
}
