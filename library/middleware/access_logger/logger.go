package access_logger

import (
	"bytes"
	"dev.taoism.gz.cn/go-taoism/util/string_tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		hostName, _ := os.Hostname()
		traceId := CreateTraceId(hostName)
		c.Set("traceId", traceId)
		c.Set("host", hostName)

		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		body := make([]byte, 0)
		// 读取body后复原
		if c.Request.Body != nil {
			body, _ = ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}

		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.String("traceId", traceId),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("body", string(body)),
			zap.String("clientIp", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("host", hostName),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func CreateTraceId(hostName string) string {

	timeNow := time.Now()
	nano, _ := strconv.Atoi(fmt.Sprintf("%d", timeNow.UnixNano())[12:16])

	traceId := fmt.Sprintf("%d%s%s", timeNow.Year()-2020, strconv.FormatInt(int64(timeNow.Month()*10000)+int64(nano), 16), strings.ReplaceAll(strings.ReplaceAll(hostName, "-", ""), ".", ""))
	if len(traceId) > 32 {
		traceId = traceId[:32]
	}

	traceId += string_tool.GetRandomStrWithNumber(32 - len(traceId))
	traceId = fmt.Sprintf("%s-%s-%s-%s", traceId[0:8], traceId[8:12], traceId[12:16], traceId[16:32])

	return strings.ToUpper(traceId)
}
