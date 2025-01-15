package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"receipt-processor/api/controllers"
	"receipt-processor/api/utils"
)

func LoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			log := make(map[string]interface{})

			log["request_id"] = params.Keys["X-Request-ID"]
			log["status_code"] = params.StatusCode
			log["path"] = params.Path
			log["method"] = params.Method
			log["start_time"] = params.TimeStamp.Format("2006/01/02 - 15:04:05")
			log["remote_addr"] = params.ClientIP
			log["response_time"] = params.Latency.String()

			s, _ := json.Marshal(log)
			return string(s) + "\n"
		},
	)
}

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var xRequestID string
		if xRequestID = c.Request.Header.Get("X-Request-ID"); xRequestID == "" {
			xRequestID = utils.GenerateUUID()
		}
		c.Set("X-Request-ID", xRequestID)
		log := logrus.WithField("request_id", xRequestID)
		c.Set("log", log)
		c.Next()
	}
}

func Run() {
	r := gin.New()
	r.Use(LoggerMiddleware())
	r.Use(RequestID())
	controllers.Initialize(r)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
