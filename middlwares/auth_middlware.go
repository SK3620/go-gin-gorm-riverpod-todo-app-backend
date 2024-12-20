package middlwares

import (
	"go-gin-gorm-riverpod-todo-app/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// gin.HandlerFunc型を返す→ type HandlerFunc func(*Context)
func AuthMiddlware(authService services.IAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization") // Authorizationヘッダー取得
		if header == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Authorizationヘッダーが、「"Bearer "」で始まることを確認
		if !strings.HasPrefix(header, "Bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 文字列 header から先頭にある "Bearer " を取り除いて、トークン部分だけを取得
		tokenString := strings.TrimPrefix(header, "Bearer ")

		// トークンに含まれる情報を基にユーザー情報を取得
		user, error := authService.GetUserFromToken(tokenString)
		if error != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return 
		}

		// ユーザー情報をリクエストコンテキストにセット（"user"キーで値を取り出す）
		ctx.Set("user", user)

		// 処理フローを次のmiddlwareまたは、目的の処理に移す
		ctx.Next()
	}
}