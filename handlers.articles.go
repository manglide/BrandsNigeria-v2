// handlers.article.go

package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	allItems, _ := getAllItems()
	itemsCount := getAllItemsCount()
	products, err := getAllItemsFrontPage()
	if err != nil {
		render(c, gin.H{"title": "Server Error", "message": http.StatusServiceUnavailable}, "500.tmpl")
	}
	loggedInInterface, _ := c.Get("is_logged_in")
	loggedIn := loggedInInterface.(bool)
	if loggedIn {
		render(c,
			gin.H{
				"title":        "Home Page",
				"payload":      articles,
				"items":        allItems,
				"itemscount":   itemsCount,
				"products":     products,
				"is_logged_in": true},
			"index.tmpl")
	} else {
		render(c,
			gin.H{
				"title":        "Home Page",
				"payload":      articles,
				"items":        allItems,
				"itemscount":   itemsCount,
				"products":     products,
				"is_logged_in": false},
			"index.tmpl")
	}

}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getSingleArticle(articleID); err == nil {
			render(c, gin.H{"title": article.Title, "message": article, "is_logged_in": true}, "article.tmpl")
		} else {
			render(c, gin.H{"title": "404 Not Found", "message": http.StatusNotFound}, "404.tmpl")
		}

	} else {
		render(c, gin.H{"title": "404 Not Found", "message": "Oops! Sorry we cant find any article with the ID"}, "404.tmpl")
	}
}

func getArticleUnAuthenticated(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getSingleArticle(articleID); err == nil {
			render(c, gin.H{"title": article.Title, "message": article}, "article.tmpl")
		} else {
			render(c, gin.H{"title": "404 Not Found", "message": http.StatusNotFound}, "404.tmpl")
		}

	} else {
		render(c, gin.H{"title": "404 Not Found", "message": "Oops! Sorry we cant find any article with the ID"}, "404.tmpl")
	}
}

func showArticleCreationPage(c *gin.Context) {
	render(c, gin.H{
		"title":        "Create New Article",
		"is_logged_in": true}, "create-article.tmpl")
}

func createArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	if a, err := createNewArticle(title, content); err == nil {
		render(c, gin.H{
			"title":        "Submission Successful",
			"payload":      a,
			"is_logged_in": true}, "submission-successful.tmpl")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

func getProductPageAuthenticated(c *gin.Context) {
	productTitle := c.Param("product_id")
	pid, err := getPID(productTitle)
	if err != nil {
		render(c, gin.H{"title": "Server Error", "message": http.StatusServiceUnavailable}, "500.tmpl")
	}
	numberofcomments := getAllCommentsCount(pid)
	data, err := getProductData(productTitle)
	if err != nil {
		render(c, gin.H{"title": "Server Error", "message": http.StatusServiceUnavailable}, "500.tmpl")
	}
	comments, err := getAllComments(pid)
	if err != nil {
		render(c, gin.H{"title": "Server Error", "message": http.StatusServiceUnavailable}, "500.tmpl")
	}
	render(c,
		gin.H{
			"title":            "Product Page - " + productTitle,
			"data":             data,
			"numberofcomments": numberofcomments,
			"comments":         comments,
			"is_logged_in":     true,
		},
		"productPage.tmpl")
}

func getProductPage(c *gin.Context) {
	productTitle := c.Param("product_id")
	pid, err := getPID(productTitle)
	if err != nil {
		log.Println(err.Error())
	}
	numberofcomments := getAllCommentsCount(pid)
	data, err := getProductData(productTitle)
	if err != nil {
		log.Println(err.Error())
	}
	comments, err := getAllComments(pid)
	if err != nil {
		// render(c, gin.H{"title": "Server Error", "message": http.StatusServiceUnavailable}, "500.tmpl")
		log.Println(err.Error())
	}
	render(c,
		gin.H{
			"title":            "Product Page - " + productTitle,
			"data":             data,
			"numberofcomments": numberofcomments,
			"comments":         comments,
		},
		"productPage.tmpl")
}

func postComments(c *gin.Context) {
	pid := c.PostForm("productid")
	cat := c.PostForm("productcategory")
	username := c.PostForm("username")
	comments := c.PostForm("comments")
	rating := c.PostForm("rating")
	log.Println(pid, cat, username, comments, rating)
}

func createProductPage(c *gin.Context) {
	categories, err := getCategories()
	if err != nil {
		render(c, gin.H{"title": "Server Error", "message": http.StatusServiceUnavailable}, "500.tmpl")
	}
	competitors, err := getCompetitors()
	if err != nil {
		render(c, gin.H{"title": "Server Error", "message": http.StatusServiceUnavailable}, "500.tmpl")
	}
	render(c,
		gin.H{
			"title":        "Create New Product",
			"is_logged_in": true,
			"competitors":  competitors,
			"categories":   categories,
		},
		"create-product.tmpl")
}

type Myform struct {
	PRODUCTGUID         string   `form:"productguid"`
	PRODUCTNAME         string   `form:"productname"`
	IMAGEDEFAULT        string   `form:"imagedefault"`
	CATEGORIES          string   `form:"categories"`
	MANUFACTURER        string   `form:"manufacturer"`
	MANUFACTURERADDRESS string   `form:"manufactureraddress"`
	ABOUT               string   `form:"about"`
	INGREDIENTS         string   `form:"ingredients"`
	PRICE               string   `form:"price"`
	COMPETITORS         []string `form:"competitors"`
	PRODUCTIMAGE1       string   `form:"productimage1"`
	PRODUCTIMAGE2       string   `json:"productimage2"`
	SKU                 string   `form:"sku"`
	MPN                 string   `form:"mpn"`
}

func genGUID(str string) string {
	c := strings.Split(str, " ")
	v := strings.Join(c, "-")
	return v
}

func createProduct(c *gin.Context) {

	var myform Myform
	myform.PRODUCTGUID = genGUID(c.PostForm("productname"))
	myform.PRODUCTNAME = c.PostForm("productname")
	myform.IMAGEDEFAULT = c.PostForm("imagedefault")
	if c.PostForm("imagedefault") == "on" {
		myform.PRODUCTIMAGE1 = "images/default-home.jpg"
		myform.PRODUCTIMAGE2 = "images/default-product.jpg"
	}
	myform.CATEGORIES = c.PostForm("categories")
	myform.MANUFACTURER = c.PostForm("manufacturer")
	myform.MANUFACTURERADDRESS = c.PostForm("manufactureraddress")
	myform.ABOUT = c.PostForm("about")
	myform.INGREDIENTS = c.PostForm("ingredients")
	myform.PRICE = c.PostForm("price")
	myform.COMPETITORS = c.Request.Form["competitors[]"]
	myform.SKU = strconv.Itoa(genSKU())
	myform.MPN = strconv.Itoa(genMPN())

	newProduct, err := makeProduct(myform)
	if err == nil {
		render(c, gin.H{
			"title":        "Submission Successful",
			"product":      newProduct,
			"is_logged_in": true}, "product-successful.tmpl")
	} else {
		render(c, gin.H{"title": "Server Error", "message": http.StatusInternalServerError}, "500.tmpl")
	}
}
