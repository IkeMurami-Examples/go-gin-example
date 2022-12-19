/*
Copyright © 2022 Ike Murami <murami.ike@gmail.com>
*/
package middleware

import (
	"context"
	"strings"
	"time"

	"github.com/IkeMurami-Examples/go-gin-example/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logger logs request/response pair.
func Logger(ctx context.Context) gin.HandlerFunc {

	logger := utils.LoggerFromContext(ctx)

	return func(ctx *gin.Context) {
		// log if log level Debug of Trace
		logLevel := logger.Level()
		if logLevel != zap.DebugLevel {
			ctx.Next()
			return
		}

		// before request

		request := ctx.Request

		var scheme string
		if request.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}

		proto := request.Proto
		method := request.Method
		remoteAddr := request.RemoteAddr
		userAgent := request.UserAgent()
		uri := strings.Join([]string{scheme, "://", request.Host, request.RequestURI}, "")

		logger.Debug(
			"Request started",
			zap.String("http-scheme", scheme),
			zap.String("http-proto", proto),
			zap.String("http-method", method),
			zap.String("remote-addr", remoteAddr),
			zap.String("user-agent", userAgent),
			zap.String("uri", uri),
		)

		timeNow := time.Now().In(time.UTC)

		// request

		ctx.Next()

		// after request

		timeNowDelta := time.Now().In(time.UTC).Sub(timeNow)

		logger.Debug(
			"Request completed",
			zap.String("http-scheme", scheme),
			zap.String("http-proto", proto),
			zap.String("http-method", method),
			zap.String("remote-addr", remoteAddr),
			zap.String("user-agent", userAgent),
			zap.String("uri", uri),
			zap.String("elapsed", timeNowDelta.String()),
		)
	}

}
