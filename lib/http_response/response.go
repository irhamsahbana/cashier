package http_response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Message    interface{} `json:"message"`
	Timestamp  string      `json:"timestamp"`
	Data       interface{} `json:"data"`
	Meta       interface{} `json:"meta,omitempty"`
}

func JSON(c *gin.Context, statusCode int, message interface{}, data interface{}, metas ...any) {
	var meta interface{}
	for _, v := range metas {
		meta = v
		break
	}

	c.JSON(statusCode, Response{
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Message:    message,
		Timestamp:  time.Now().Format(time.RFC3339Nano),
		Data:       data,
		Meta:       meta,
	})
}
