package middleware

import (
	"net/http"
	"strings"

	"WebApp/internal/app/auth"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware проверяет JWT из заголовка Authorization или cookie
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tokenStr string

		// 1. Пробуем получить из заголовка: "Bearer <token>"
		authHeader := ctx.GetHeader("Authorization")
		if authHeader != "" {
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) == 2 && parts[0] == "Bearer" {
				tokenStr = parts[1]
			}
		}

		// 2. Если нет — пробуем из cookie
		if tokenStr == "" {
			cookie, err := ctx.Cookie("token")
			if err == nil {
				tokenStr = cookie
			}
		}

		if tokenStr == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization required"})
			return
		}

		claims, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// Сохраняем данные пользователя в контексте для хендлеров
		ctx.Set("user_id", claims.UserID)
		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}

// ModeratorOnly разрешает доступ только пользователям с ролью "moderator"
func ModeratorOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get("role")
		if !exists || role != "moderator" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "moderator access required"})
			return
		}
		ctx.Next()
	}
}
