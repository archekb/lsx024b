package middleware

import (
	"time"

	"github.com/archekb/lsx024b/internal/log"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	httpLog := log.StandartNamed("HTTP Server")
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		end := time.Now()

		latency := end.Sub(start)
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		httpLog.Infof("%3d | %13v | %15s | %s | %s ", statusCode, latency, clientIP, method, path)
	}
}
