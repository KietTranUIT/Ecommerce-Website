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
		token, err := c.Cookie("admin-token")

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

		result, _, username := util.VerifyToken(token)
		c.Set("userId", username)
		if !result {
			c.Redirect(302, "/admin/login")
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("user-token")

		if err != nil {
			if strings.Contains(err.Error(), "http: named cookie not present") {
				c.Redirect(302, "/login")
				c.Abort()
			}
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Server Error"})
			c.Abort()
			return
		}

		if token == "" {
			c.Redirect(302, "/login")
			c.Abort()
			return
		}

		result, _, username := util.VerifyToken(token)
		c.Set("userId", username)
		if !result {
			c.Redirect(302, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
