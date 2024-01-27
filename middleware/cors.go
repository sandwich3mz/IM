package middleware

import (
	"github.com/gin-contrib/cors"
)

// GetCorsConfig 获取第三方宝解决跨域问题的初始化的配置，并自定义
func GetCorsConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowWildcard = true
	corsConfig.AllowOrigins = []string{"http://localhost:*", "http://110.42.208.9:*"}
	addedHeaders := []string{"Access-Token", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "Cache-Control", "X-Requested-With", "Hmac-Key", "URI", "Hmac", "Hashed-Secret", "Secret"}
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, addedHeaders...)
	corsConfig.AllowCredentials = true
	return corsConfig
}
