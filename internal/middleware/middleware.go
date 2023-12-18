package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticateAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Xin chao ban")
		token, err := c.Cookie("bear")

		if err != nil {
			if strings.Contains(err.Error(), "http: named cookie not present") {
				c.Redirect(301, "/admin/login")
			}
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Server Error"})
			return
		}

		if token == "" {
			log.Println("OK")
			c.Redirect(301, "/admin/login")
			return
		}

		c.Next()
	}
}
