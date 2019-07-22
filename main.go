// main.go

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/brandsnigeria/webapp/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var router *gin.Engine
var c *gin.Context

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

func add(index int) int {
	return index + 1
}

func removeNewLines(item string) string {
	return strings.ReplaceAll(item, "\n", "")
}

func splitUtil(item string) []string {
	return strings.Split(item, "~,")
}

func splitRLength(item string) int {
	var c []string
	c = strings.Split(item, "~,")
	return len(c)
}

func visualDate(val string) string {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, val)
	if err != nil {
		log.Println(err)
	}
	return t.Format(time.RFC1123)
}

func trendDir(val string) string {
	if val == "Up" {
		return "btn btn-success btn-lg"
	} else {
		return "btn btn-danger btn-lg"
	}
}

func trendDirT(val string) string {
	if val == "Up" {
		return "Good Uptrend"
	} else {
		return "Bad Downtrend"
	}
}

func sentimentMood(val string) string {
	if val == "Good" {
		return "btn btn-success btn-lg"
	} else {
		return "btn btn-danger btn-lg"
	}
}

func sentimentMoodT(val string) string {
	if val == "Good" {
		return "Good Sentiment"
	} else {
		return "Bad Sentiment"
	}
}

func splitRDateIndex(item string, index int) string {
	var q []string
	var seconds int
	q = strings.Split(item, "#")
	seconds, _ = strconv.Atoi(q[index])
	seconds = seconds
	return time.Unix(int64(seconds), 0).Format(time.RFC850)
}

func splitRIndex(item string, index int) string {
	var u []string
	u = strings.Split(item, "#")
	return u[index]
}

func uppercase(item string) string {
	return strings.ToUpper(item)
}

func iterate(count string) []int {
	s, err := strconv.ParseFloat(count, 64)
	if err != nil {
		log.Println(err)
	}
	var i int
	var Items []int
	for i = 0; i < (int(s)); i++ {
		Items = append(Items, i)
	}
	return Items
}

func equal(val, val2 int) bool {
	if add(val) == val2 {
		return true
	} else {
		return false
	}
}

func roundUp(val string) string {
	i, _ := strconv.ParseFloat(val, 0)
	s := fmt.Sprintf("%.2f", i)
	return s
}
func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.SetFuncMap(template.FuncMap{
		"add":             add,
		"equal":           equal,
		"removeNewLines":  removeNewLines,
		"uppercase":       uppercase,
		"iterate":         iterate,
		"splitUtil":       splitUtil,
		"splitRIndex":     splitRIndex,
		"splitRLength":    splitRLength,
		"splitRDateIndex": splitRDateIndex,
		"visualDate":      visualDate,
		"trendDir":        trendDir,
		"trendDirT":       trendDirT,
		"sentimentMood":   sentimentMood,
		"sentimentMoodT":  sentimentMoodT,
		"roundUp":         roundUp,
	})
	router.LoadHTMLGlob("templates/*.*")
	router.Static("/css", "templates/css")
	router.Static("/js", "templates/js")
	router.Static("/vendor", "templates/vendor")

	db, err := sql.Open("mysql", "reviewmonster:love~San&500#@tcp(127.0.0.1:3306)/asknigeria?charset=utf8mb4,utf8")
	if err != nil {
		render(c, gin.H{"title": "Server Error", "message": http.StatusServiceUnavailable}, "500.tmpl")
	}
	defer db.Close()

	database.DB = db // Initialize the routes

	initializeRoutes()

	// Start serving the application
	router.Run(":9999")

}
