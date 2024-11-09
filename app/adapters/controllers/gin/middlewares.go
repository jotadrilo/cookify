package gin

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/jotadrilo/cookify/internal/logger"
)

var ginlog = logger.Default().WithOptions(zap.WithCaller(false)).Named("GIN")

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Other middlewares may modify these fields
		var (
			start = time.Now().UTC()
			path  = c.Request.URL.Path
			query = c.Request.URL.RawQuery
		)

		c.Next()

		var (
			status  = c.Writer.Status()
			end     = time.Now().UTC()
			latency = end.Sub(start)
		)

		fields := []zapcore.Field{
			zap.Int("status", status),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
			zap.String("time", end.Format(time.RFC3339)),
		}

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				ginlog.Error(e, fields...)
			}
		}

		if status <= http.StatusOK {
			ginlog.Info(path, fields...)
		} else {
			ginlog.Warn(path, fields...)
		}
	}
}

func ZapRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					ginlog.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) //nolint: errcheck
					c.Abort()
					return
				}

				ginlog.Error("[Recovery from panic]",
					zap.Time("time", time.Now()),
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.String("stack", string(debug.Stack())),
				)

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}
