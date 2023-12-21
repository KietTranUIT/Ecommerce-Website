package middleware

import (
	"log"
	"net/http"
	"strings"
	"user-service/internal/core/common/util"

	"github.com/gin-gonic/gin"
)

func AuthenticateAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Hello")
		token, err := c.Cookie("bear")

		if err != nil {
			if strings.Contains(err.Error(), "http: named cookie not present") {
				log.Println("3")
				c.Redirect(302, "/admin/login")
				c.Abort()
			}
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Server Error"})
			c.Abort()
			return
		}

		if token == "" {
			log.Println("2")
			c.Redirect(302, "/admin/login")
			c.Abort()
			return
		}

		result, _ := util.VerifyToken(token)

		if !result {
			log.Println("1")
			c.Redirect(302, "/admin/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
