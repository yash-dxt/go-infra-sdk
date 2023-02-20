package config

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/metaphi-org/go-infra-sdk/constants"
)

func CreateGinCorsConfig() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AddAllowHeaders(constants.AuthHeaderKey)
	config.AllowCredentials = true
	config.AllowOriginFunc = func(origin string) bool {
		return strings.HasSuffix(origin, ".awen.finance") ||
			origin == "http://localhost:3000"
	}

	return cors.New(config)
}
