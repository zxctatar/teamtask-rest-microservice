package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSessionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("sessionId")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "needed cookie with sessionId",
			})
			ctx.Abort()
			return
		}
		ctx.Set("sessionId", cookie)
		ctx.Next()
	}
}
