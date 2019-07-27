// middleware.auth.go
package main

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ensureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.HTML(http.StatusUnauthorized, "unauthenticated.tmpl", gin.H{
				"ErrorTitle":   "Unauthorized Access",
				"is_logged_in": false,
				"ErrorMessage": "Unauthorised Access"})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func ensureLoggedInJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if tokenString, err := c.Cookie("token"); err == nil || tokenString != "" {
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					// c.AbortWithStatus(http.StatusUnauthorized)
					c.HTML(http.StatusUnauthorized, "unauthenticated.tmpl", gin.H{
						"ErrorTitle":   "Unauthorized Access",
						"is_logged_in": false,
						"ErrorMessage": err.Error()})
				}
				// hmacSampleSecret is a []byte containing your secret, e.g.

				return secretKey, nil
			})
			if err != nil {
				// c.AbortWithStatus(http.StatusUnauthorized)
				c.HTML(http.StatusUnauthorized, "unauthenticated.tmpl", gin.H{
					"ErrorTitle":   "Unauthorized Access",
					"is_logged_in": false,
					"ErrorMessage": err.Error()})
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// If token is valid
				c.Next()
			} else {
				// c.AbortWithStatus(http.StatusUnauthorized)
				c.HTML(http.StatusUnauthorized, "unauthenticated.tmpl", gin.H{
					"ErrorTitle":   "Unauthorized Access",
					"is_logged_in": false,
					"ErrorMessage": err.Error()})
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			// c.AbortWithStatus(http.StatusUnauthorized)
			c.HTML(http.StatusUnauthorized, "unauthenticated.tmpl", gin.H{
				"ErrorTitle":   "Unauthorized Access",
				"is_logged_in": false,
				"ErrorMessage": err.Error()})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func ensureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func neutral() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			c.Next()
		} else {
			c.Next()
		}
	}
}

func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}
