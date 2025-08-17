package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "X-Request-ID"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader(RequestIDKey)
		if id == "" { id = uuid.NewString() }
		c.Header(RequestIDKey, id)
		c.Set(RequestIDKey, id)
		c.Next()
	}
}
