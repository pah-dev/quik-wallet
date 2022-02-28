package middlewares

import (
	"github.com/gin-gonic/gin"
)

const (
	XRequestIdKey = "X-Request-ID" // request id header key
)

// RequestIDMiddleware attach request id and logger to context
// 1. extract request id from header if exist, otherwise generate
// 2. attach request id to logger and store it to context
func RequestIDMiddleware() gin.HandlerFunc {
	// return func(c *gin.Context) {
	// 	requestId := c.Request.Header.Get(XRequestIdKey)
	// 	if requestId == "" {
	// 		requestId = uuid.New().String()
	// 	}

	// 	c.Request = c.Request.WithContext(trace.WithRequestID(c, requestId))
	// 	logger := logging.DefaultLogger().With("requestId", requestId)
	// 	c.Request = c.Request.WithContext(logging.WithLogger(c, logger))
	// 	c.Writer.Header().Set(XRequestIdKey, requestId)
	// }
	return nil
}