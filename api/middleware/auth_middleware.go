package middleware

import (
	"brodo-demo/service/security"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func authMiddleware(tokenManager security.AuthenticationTokenManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "

		authValue := ctx.GetHeader("Authorization")

		if len(authValue) < len(BEARER_SCHEMA) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authValue, BEARER_SCHEMA)

		userId, err := tokenManager.VerifyAccessToken(token)

		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("userId", userId)
		ctx.Next()
	}
}

func NewAuthMiddleware(tokenManager security.AuthenticationTokenManager) gin.HandlerFunc {
	return authMiddleware(tokenManager)
}
