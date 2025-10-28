package middleware

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func AuthJwt() http.Middleware {
	return func(ctx http.Context) {
		// Check jwt token validity here
		token := ctx.Request().Header("Authorization")

		// If no token is provided, abort with unauthorized status
		if token == "" {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"message": "Unauthorized: No token provided",
			})
		}

		// Validate the token (this is a placeholder, implement your own logic)
		_, err := facades.Auth(ctx).Parse(token)

		// If token is invalid, abort with unauthorized status
		if err != nil {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"message": "Unauthorized: Invalid token",
			})
		}

		// Token is valid, proceed to the next middleware/handler
		ctx.Request().Next()
	}
}
