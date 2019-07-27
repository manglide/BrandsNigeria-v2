// handlers.user.go
package main

import (
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte(os.Getenv("SESSION_SECRET"))
var j *gin.Context

func generateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

func showRegistrationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Register"}, "register.tmpl")
}

func register(c *gin.Context) {
	// Obtain the POSTed username and password values
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	username := c.PostForm("username")
	password := c.PostForm("password")
	if user, err := registerNewUser(firstname, lastname, username, password); err == nil {
		// If the user is created, set the token in a cookie and log the user in
		// token, err := generateSessionTokenJWT(username)
		token := generateSessionToken()
		// if err != nil {
		// 	c.HTML(http.StatusBadRequest, "login.html", gin.H{
		// 		"ErrorTitle":   "Failed to Generate JWT Token",
		// 		"ErrorMessage": "Failed to generate JWT Token " + err.Error()})
		// }
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		render(c, gin.H{
			"title":        "Successful registration & Login",
			"user":         &user,
			"is_logged_in": true}, "registeration-successful.tmpl")
	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest, "register.tmpl", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})
	}
}

func showLoginPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Login",
	}, "login.tmpl")
}

func showAboutPage(c *gin.Context) {
	render(c, gin.H{
		"title": "About BrandsNigeria",
	}, "about.tmpl")
}

func showAboutPageAuthenticated(c *gin.Context) {
	if Superadmin > 0 {
		render(c, gin.H{
			"title":        "About BrandsNigeria",
			"is_logged_in": true,
			"superadmin":   Superadmin,
		}, "about.tmpl")
	} else {
		render(c, gin.H{
			"title":        "About BrandsNigeria",
			"is_logged_in": true,
			"superadmin":   Superadmin,
		}, "about.tmpl")
	}

}

func showFeedbackPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Feedback",
	}, "feedback.tmpl")
}

func showFeedbackPageAuthenticated(c *gin.Context) {
	if Superadmin > 0 {
		render(c, gin.H{
			"title":        "Feedback",
			"is_logged_in": true,
			"superadmin":   Superadmin,
		}, "feedback.tmpl")
	} else {
		render(c, gin.H{
			"title":        "Feedback",
			"is_logged_in": true,
			"superadmin":   Superadmin,
		}, "feedback.tmpl")
	}
}

type User struct {
	username string `json:"username"`
}

func performLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if isUserValid(username, password) {
		token, err := generateSessionTokenJWT(username)
		if err != nil {
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
				"ErrorTitle":   "Failed to Generate JWT Token",
				"ErrorMessage": "Failed to generate JWT Token " + err.Error()})
		}
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		if username == "08146382332" {
			Superadmin = 1
			UserLoggedIn = username
			// render(c, gin.H{
			// 	"title":        "Successful Login",
			// 	"user":         username,
			// 	"is_logged_in": true, "superadmin": "08146382332"}, "login-successful.tmpl")
			showIndexPage(c)
		} else {
			Superadmin = 0
			UserLoggedIn = username
			showIndexPage(c)
			// render(c, gin.H{
			// 	"title":        "Successful Login",
			// 	"user":         username,
			// 	"is_logged_in": true, "superadmin": "no"}, "login-successful.tmpl")
		}

	} else {
		c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Set("is_logged_in", false)
	c.Set("username", nil)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func generateSessionTokenJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username":  username,
		"ExpiresAt": 15000,
		"IssuedAt":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
