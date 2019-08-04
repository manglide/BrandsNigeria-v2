// handlers.article.go

package main

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	// articles := getAllArticles()
	allItems, _ := getAllItems()
	itemsCount := getAllItemsCount()
	products, err := getAllItemsFrontPage()
	if err != nil {
		render(c, gin.H{"title": "Server Error", "message": http.StatusServiceUnavailable}, "500.tmpl")
	}
	loggedInInterface, _ := c.Get("is_logged_in")
	loggedIn := loggedInInterface.(bool)

	if loggedIn {
		if Superadmin > 0 {
			render(c,
				gin.H{
					"title":        "Home Page",
					"items":        allItems,
					"itemscount":   itemsCount,
					"products":     products,
					"superadmin":   Superadmin,
					"username":     UserLoggedIn,
					"is_logged_in": true},
				"index.tmpl")
		} else {
			render(c,
				gin.H{
					"title":        "Home Page",
					"items":        allItems,
					"itemscount":   itemsCount,
					"products":     products,
					"superadmin":   Superadmin,
					"username":     UserLoggedIn,
					"is_logged_in": true},
				"index.tmpl")
		}
	} else {
		render(c,
			gin.H{
				"title":        "Home Page",
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
			"username":         UserLoggedIn,
			"superadmin":       Superadmin,
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
			"username":         UserLoggedIn,
			"superadmin":       Superadmin,
		},
		"productPage.tmpl")
}

func postComments(c *gin.Context) {
	pid := c.PostForm("productid")
	cat := c.PostForm("productcategory")
	username := UserLoggedIn
	author := c.PostForm("author")
	comments := c.PostForm("comment")
	rating := c.PostForm("rating")
	sentiment := c.PostForm("sentiment")
	latitude := c.PostForm("latitude")
	longitude := c.PostForm("longitude")
	_, err := insertComments(pid, cat, username, comments,
		rating, sentiment, latitude, longitude, author)
	if err != nil {
		c.JSON(400, gin.H{
			"data":    "failed",
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"data":    "success",
			"message": "reload",
		})
	}
}

func createProductPage(c *gin.Context) {
	categories, err := getCategories()
	if err != nil {
		render(c, gin.H{"title": "Server Error", "message": http.StatusServiceUnavailable}, "500.tmpl")
	}
	competitors, err := getCompetitorsI()
	if err != nil {
		render(c, gin.H{"title": "Server Error", "message": http.StatusServiceUnavailable}, "500.tmpl")
	}
	render(c,
		gin.H{
			"title":        "Create New Product",
			"is_logged_in": true,
			"superadmin":   Superadmin,
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

type MyformE struct {
	PRODUCTID           string   `form:"productid"`
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
	k := strings.ToLower(str)
	c := strings.Split(k, " ")
	v := strings.Join(c, "-")
	return v
}

func createProduct(c *gin.Context) {

	var myform Myform
	myform.PRODUCTGUID = genGUID(c.PostForm("productname"))
	myform.PRODUCTNAME = c.PostForm("productname")
	myform.IMAGEDEFAULT = c.PostForm("imagedefault")

	if c.PostForm("imagedefault") == "yes" {
		myform.PRODUCTIMAGE1 = "images/default-home.jpg"
		myform.PRODUCTIMAGE2 = "images/default-product.jpg"
	} else {

		// file, err := c.FormFile("imagehomepage")
		// if err != nil {
		// 	log.Println(err)
		// }

		// err = c.SaveUploadedFile(file, "templates/assets/images/"+genGUID(c.PostForm("productname")))
		// if err != nil {
		// 	log.Println(err)
		// }

		// fileX, errX := c.FormFile("imagemain")
		// if errX != nil {
		// 	log.Println(errX)
		// }

		// errX = c.SaveUploadedFile(fileX, "templates/assets/images/"+fileX.Filename)
		// if errX != nil {
		// 	log.Println(errX)
		// }

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

func reviewLikes(c *gin.Context) {
	p1, _ := url.QueryUnescape(c.PostForm("data1"))
	p2, _ := url.QueryUnescape(c.PostForm("data2"))
	p3, _ := url.QueryUnescape(c.PostForm("data3"))
	q := []string{p1, p2, p3}
	list, err := getReviewLikes(q)
	if err != nil {
		// render(c, gin.H{"title": "Server Error", "message": http.StatusInternalServerError}, "500.tmpl")
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"data":    list,
		"message": "success",
	})
}

func reviewDisLikes(c *gin.Context) {
	p1, _ := url.QueryUnescape(c.PostForm("data1"))
	p2, _ := url.QueryUnescape(c.PostForm("data2"))
	p3, _ := url.QueryUnescape(c.PostForm("data3"))
	q := []string{p1, p2, p3}
	listD, err := getReviewDisLikes(q)
	if err != nil {
		// render(c, gin.H{"title": "Server Error", "message": http.StatusInternalServerError}, "500.tmpl")
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"data":    listD,
		"message": "success",
	})
}

func reviewRatings(c *gin.Context) {
	p1, _ := url.QueryUnescape(c.PostForm("data1"))
	p2, _ := url.QueryUnescape(c.PostForm("data2"))
	p3, _ := url.QueryUnescape(c.PostForm("data3"))
	q := []string{p1, p2, p3}
	listZ, err := getReviewRating(q)
	if err != nil {
		// render(c, gin.H{"title": "Server Error", "message": http.StatusInternalServerError}, "500.tmpl")
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"data":    listZ,
		"message": "success",
	})
}

func getAreasOfAcceptance(c *gin.Context) {
	p1, _ := url.QueryUnescape(c.PostForm("data"))
	listZ, err := getAcceptanceAreas(p1)
	if err != nil {
		// render(c, gin.H{"title": "Server Error", "message": http.StatusInternalServerError}, "500.tmpl")
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"data":    listZ,
		"message": "success",
	})
}

func getAreasOfRejection(c *gin.Context) {
	p1, _ := url.QueryUnescape(c.PostForm("data"))
	listZ, err := getRejectionAreas(p1)
	if err != nil {
		// render(c, gin.H{"title": "Server Error", "message": http.StatusInternalServerError}, "500.tmpl")
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"data":    listZ,
		"message": "success",
	})
}

func getProductRecommendation(c *gin.Context) {
	p1, _ := url.QueryUnescape(c.PostForm("data1"))
	p2, _ := url.QueryUnescape(c.PostForm("data2"))
	p3, _ := url.QueryUnescape(c.PostForm("data3"))
	q := []string{genGUID(p1), genGUID(p2), genGUID(p3)}
	listZ, listY, err := productRecommendation(q)
	if err != nil {
		// render(c, gin.H{"title": "Server Error", "message": http.StatusInternalServerError}, "500.tmpl")
		log.Println(err)
	}

	loggedInInterface, _ := c.Get("is_logged_in")
	loggedIn := loggedInInterface.(bool)
	if loggedIn {
		if len(listZ) > 0 && len(listY) == 0 {
			render(c,
				gin.H{
					"products":     listZ,
					"is_logged_in": true},
				"product-recommendation-card-auth.tmpl")
		} else if len(listZ) == 0 && len(listY) > 0 {
			render(c,
				gin.H{
					"products":     listY,
					"is_logged_in": true},
				"product-no-competition-auth.tmpl")
		}
	} else {
		if len(listZ) > 0 && len(listY) == 0 {
			render(c,
				gin.H{
					"products":     listZ,
					"is_logged_in": true},
				"product-recommendation-card.tmpl")
		} else if len(listZ) == 0 && len(listY) > 0 {
			render(c,
				gin.H{
					"products":     listY,
					"is_logged_in": true},
				"product-no-competition.tmpl")
		}
	}

}

func pCompetitor(c *gin.Context) {
	v, _ := url.QueryUnescape(c.PostForm("data"))
	v = genGUID(v)
	listZ, listY, err := getCompetitors(v)
	if err != nil {
		// render(c, gin.H{"title": "Server Error", "message": http.StatusInternalServerError}, "500.tmpl")
		log.Println(err)
	}
	loggedInInterface, _ := c.Get("is_logged_in")
	loggedIn := loggedInInterface.(bool)
	if loggedIn {
		if len(listZ) > 0 && len(listY) == 0 {
			render(c,
				gin.H{
					"products":     listZ,
					"is_logged_in": true},
				"product-competitor-auth.tmpl")
		} else if len(listZ) == 0 && len(listY) > 0 {
			render(c,
				gin.H{
					"products":     listY,
					"is_logged_in": true},
				"product-no-competition-auth.tmpl")
		}
	} else {
		if len(listZ) > 0 && len(listY) == 0 {
			render(c,
				gin.H{
					"products":     listZ,
					"is_logged_in": true},
				"product-competitor.tmpl")
		} else if len(listZ) == 0 && len(listY) > 0 {
			render(c,
				gin.H{
					"products":     listY,
					"is_logged_in": true},
				"product-no-competition.tmpl")
		}
	}
}

func createProductListPage(c *gin.Context) {
	v, err := getProductLists()
	if err != nil {
		// render(c, gin.H{"title": "Server Error", "message": http.StatusInternalServerError}, "500.tmpl")
		log.Println(err)
	}
	render(c, gin.H{
		"title":        "Submission Successful",
		"data":         v,
		"is_logged_in": true, "superadmin": true}, "product-list.tmpl")
}

func createDeletedProductListPage(c *gin.Context) {
	v, err := getDeletedProductLists()
	if err != nil {
		// render(c, gin.H{"title": "Server Error", "message": http.StatusInternalServerError}, "500.tmpl")
		log.Println(err)
	}
	render(c, gin.H{
		"title": "Submission Successful",
		"data":  v, "username": UserLoggedIn,
		"is_logged_in": true, "superadmin": Superadmin}, "delete-product-list.tmpl")
}

func ratedProducts(c *gin.Context) {
	v, err := getProductListsByUser(UserLoggedIn)
	if err != nil {
		// render(c, gin.H{"title": "Server Error", "message": http.StatusInternalServerError}, "500.tmpl")
		log.Println(err)
	}
	render(c, gin.H{
		"title":        "Submission Successful",
		"data":         v,
		"username":     UserLoggedIn,
		"is_logged_in": true, "superadmin": Superadmin}, "rated-products.tmpl")
}

func withdrawRating(c *gin.Context) {
	rid := c.PostForm("rid")
	pid := c.PostForm("pid")
	user := c.PostForm("username")
	a, err := deleteRating(rid, pid, user)
	if err != nil {
		c.JSON(400, gin.H{
			"data":    "failed",
			"message": err.Error(),
		})
	} else {
		if a > 0 {
			c.JSON(200, gin.H{
				"data":    "success",
				"message": "reload",
			})
		} else {
			c.JSON(200, gin.H{
				"data":    "failed",
				"message": "Item not deleted",
			})
		}
	}
}

func editProduct(c *gin.Context) {
	p := c.Param("product_id")
	categories, err := getCategories()
	if err != nil {
		render(c, gin.H{"title": "Server Error", "message": http.StatusServiceUnavailable}, "500.tmpl")
	}
	competitors, err := getCompetitorsI()
	if err != nil {
		render(c, gin.H{"title": "Server Error", "message": http.StatusServiceUnavailable}, "500.tmpl")
	}
	log.Println(p)
	data, err := getProductToEdit(p)
	log.Println(data)
	if err != nil {
		render(c, gin.H{"title": "Server Error", "message": http.StatusServiceUnavailable}, "500.tmpl")
	}
	render(c,
		gin.H{
			"title":        "Edit Product",
			"is_logged_in": true,
			"superadmin":   Superadmin,
			"competitors":  competitors,
			"categories":   categories,
			"data":         data,
		},
		"edit-product.tmpl")
}

func saveProduct(c *gin.Context) {
	var myform MyformE
	myform.PRODUCTID = c.PostForm("pid")
	myform.PRODUCTGUID = c.PostForm("productname")
	myform.PRODUCTNAME = c.PostForm("productname")
	myform.IMAGEDEFAULT = c.PostForm("imagedefault")
	if c.PostForm("imagedefault") == "yes" {
		myform.PRODUCTIMAGE1 = "images/default-home.jpg"
		myform.PRODUCTIMAGE2 = "images/default-product.jpg"
	} else {
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

	editedProduct, err := editProductDB(myform)
	if err == nil {
		render(c, gin.H{
			"title":        "Submission Successful",
			"product":      editedProduct,
			"username":     UserLoggedIn,
			"superadmin":   Superadmin,
			"is_logged_in": true}, "edit-product-successful.tmpl")
	} else {
		render(c, gin.H{"title": "Server Error", "message": http.StatusInternalServerError}, "500.tmpl")
	}
}

func deleteProduct(c *gin.Context) {
	pid := c.PostForm("pid")
	guid := c.PostForm("guid")
	b, err := deleteITEM(guid, pid)
	if err != nil {
		log.Println(err)
		render(c, gin.H{"title": "Server Error", "message": http.StatusInternalServerError}, "500.tmpl")
	}
	if b > 0 {
		c.JSON(200, gin.H{
			"data":    "success",
			"message": "reload",
		})
	} else {
		c.JSON(400, gin.H{
			"data":    "failed",
			"message": "cannot delete item",
		})
	}
}

func restoreProduct(c *gin.Context) {
	pid := c.PostForm("pid")
	guid := c.PostForm("guid")
	b, err := restoreITEM(guid, pid)
	if err != nil {
		log.Println(err)
		render(c, gin.H{"title": "Server Error", "message": http.StatusInternalServerError}, "500.tmpl")
	}
	if b > 0 {
		c.JSON(200, gin.H{
			"data":    "success",
			"message": "reload",
		})
	} else {
		c.JSON(400, gin.H{
			"data":    "failed",
			"message": "cannot restore item",
		})
	}
}

func genSitemap(c *gin.Context) {
	sitemapData, err := dataSitemap()
	if err == nil {
		// c.Writer.Header().Add("Content-Type", "application/xml; charset=utf-8")
		c.Writer.Header().Set("Content-Type", "application/xml; charset=utf-8")
		// c.Header("Content-Type", "text/xml")

		render(c, gin.H{
			"title":        "Sitemap",
			"sitemap":      sitemapData,
			"username":     UserLoggedIn,
			"superadmin":   Superadmin,
			"is_logged_in": true}, "sitemap.tmpl")
	} else {
		render(c, gin.H{"title": "Server Error", "message": http.StatusInternalServerError}, "500.tmpl")
	}
}

func approveRating(c *gin.Context) {
	reviewID := c.PostForm("reviewid")
	productID := c.PostForm("pid")
	user := c.PostForm("user")
	count, err := approveRatingDB(reviewID, productID, user)

	if count > 0 {
		c.JSON(200, gin.H{
			"data": "success",
		})
	} else {
		c.JSON(400, gin.H{
			"data": err.Error(),
		})
	}
}

func disapproveRating(c *gin.Context) {
	reviewID := c.PostForm("reviewid")
	productID := c.PostForm("pid")
	user := c.PostForm("user")
	count, err := disapproveRatingDB(reviewID, productID, user)

	if count > 0 {
		c.JSON(200, gin.H{
			"data": "success",
		})
	} else {
		c.JSON(400, gin.H{
			"data": err.Error(),
		})
	}
}
