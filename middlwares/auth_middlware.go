package middlwares

import (
	"go-gin-gorm-riverpod-todo-app/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddlware(authService services.IAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")
		user, error := authService.GetUserFromToken(tokenString)
		if error != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return 
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}