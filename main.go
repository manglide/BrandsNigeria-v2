// main.go

package main

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/brandsnigeria/webapp/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var router *gin.Engine

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}

func main() {

	db, err := sql.Open("mysql", "reviewmonster:love~San&500#@tcp(127.0.0.1:3306)/asknigeria?charset=utf8mb4,utf8")
	if err != nil {
		errors.New("Database Connection Failed")
	}
	defer db.Close()
	database.DB = db
	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/css", "templates/css")
	router.Static("/js", "templates/js")
	router.Static("/vendor", "templates/vendor")

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run(":9999")

}
