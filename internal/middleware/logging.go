package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/manattan/mumbai/internal/pkg/logger"
)

func Logging(logger logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			req := c.Request()
			res := c.Response()

			logger.Info(
				"method=%s uri=%s status=%d latency=%v remote_ip=%s user_agent=%s request_id=%s",
				req.Method,
				req.RequestURI,
				res.Status,
				time.Since(start),
				c.RealIP(),
				req.UserAgent(),
				c.Get(RequestIDHeader),
			)

			return err
		}
	}
}
