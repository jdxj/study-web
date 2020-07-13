package routers

import (
	"net/http"

	"github.com/jdxj/study-web/database"
	"github.com/jdxj/study-web/models"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
)

func RecordUserAgent() gin.HandlerFunc {
	return func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		if userAgent == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "User-Agent not found",
			})
			return
		}

		ua := &models.UserAgent{
			Value: userAgent,
		}
		if err := database.AddUserAgent(ua); err != nil {
			logs.Error("add user agent failed: %s", err)
		}
	}
}
