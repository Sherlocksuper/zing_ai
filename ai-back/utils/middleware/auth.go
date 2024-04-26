package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func AuthTokenCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		if path == "/api/user/login" || path == "/api/user/register" {
			context.Next()
		} else {
			token := context.GetHeader("Authorization")
			if checkToken(token) {
				log.Error().Msg("token is required")
				context.JSON(401, gin.H{
					"message": "token is required",
				})
				context.Abort()
			} else {
				context.Next()
			}
		}
	}
}

func checkToken(token string) bool {
	if token == "" {
		log.Error().Msg("token is required")
		return false
	}
	return true
}
