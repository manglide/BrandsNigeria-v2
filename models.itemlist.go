// models.article.go

package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"database/sql"
	"math/rand"
	"time"

	"github.com/brandsnigeria/webapp/database"
)

type itemList struct {
	PRODUCTID   int    `json:"id"`
	PRODUCTGUID string `json:"productGUID"`
}

type itemProductList struct {
	PRODUCTID           int    `json:"id"`
	PRODUCTTITLE        string `json:"productTITLE"`
	PRODUCTGUID         string `json:"productGUID"`
	PRODUCTCATEGORY     string `json:"productCATEGORY"`
	PRODUCTMANUFACTURER string `json:"productMANUFACTURER"`
}

type userRatedProductList struct {
	PRODUCTID           int    `json:"id"`
	REVIEWID            int    `json:"reviewID"`
	PRODUCTTITLE        string `json:"productTITLE"`
	RATING              string `json:"rating"`
	PRODUCTGUID         string `json:"productGUID"`
	PRODUCTCATEGORY     string `json:"productCATEGORY"`
	PRODUCTMANUFACTURER string `json:"productMANUFACTURER"`
}

type productList struct {
	PRODUCTID             int    `json:"productID"`
	PRODUCTNAME           string `json:"productNAME"`
	PRODUCTGUID           string `json:"productGUID"`
	PRODUCTDESCRIPTION    string `json:"productDESCRIPTION"`
	PRODUCTMANUFACTURER   string `json:"productMANUFACTURER"`
	PRODUCTLIKES          string `json:"productLIKES"`
	PRODUCTDISLIKES       string `json:"productDISLIKES"`
	PRODUCTTREND          string `json:"productTREND"`
	PRODUCTTRENDDIRECTION string `json:"productTRENDDIRECTION"`
	PRODUCTSENTIMENT      string `json:"productSENTIMENT"`
	PRODUCTSENTIMENTMOOD  string `json:"productSENTIMENTMOOD"`
	PRODUCTUSERCOMMENTS   string `json:"productUSERCOMMENTS"`
	PRODUCTRATING         string `json:"productRATING"`
	PRODUCTDATEPUBLISHED  string `json:"productDATEPUBLISHED"`
	PRODUCTPRICE          string `json:"productPRICE"`
	PRODUCTLOCATIONCOUNT  string `json:"productLOCATIONCOUNT"`
	PRODUCTINGREDIENTS    string `json:"productINGREDIENTS"`
	PRODUCTCATEGORY       string `json:"productCATEGORY"`
	PRODUCTIMAGE1         string `json:"productIMAGE1"`
	PRODUCTIMAGE2         string `json:"productIMAGE2"`
}

type productListPage struct {
	PRODUCTID                int    `json:"productID"`
	PRODUCTNAME              string `json:"productNAME"`
	PRODUCTGUID              string `json:"productGUID"`
	PRODUCTDESCRIPTION       string `json:"productDESCRIPTION"`
	PRODUCTMANUFACTURER      string `json:"productMANUFACTURER"`
	PRODUCTLIKES             string `json:"productLIKES"`
	PRODUCTDISLIKES          string `json:"productDISLIKES"`
	PRODUCTTREND             string `json:"productTREND"`
	PRODUCTTRENDDIRECTION    string `json:"productTRENDDIRECTION"`
	PRODUCTSENTIMENT         string `json:"productSENTIMENT"`
	PRODUCTSENTIMENTMOOD     string `json:"productSENTIMENTMOOD"`
	PRODUCTUSERCOMMENTS      string `json:"productUSERCOMMENTS"`
	PRODUCTRATING            string `json:"productRATING"`
	PRODUCTAUTHOR            string `json:"productAUTHOR"`
	PRODUCTDATEPUBLISHED     string `json:"productDATEPUBLISHED"`
	PRODUCTSKU               string `json:"productSKU"`
	PRODUCTMPN               string `json:"productMPN"`
	PRODUCTPRICE             string `json:"productPRICE"`
	PRODUCTLOCATIONCOUNT     string `json:"productLOCATIONCOUNT"`
	PRODUCTFIRSTCOMPETITION  string `json:"productFIRSTCOMPETITION"`
	PRODUCTSECONDCOMPETITION string `json:"productSECONDCOMPETITION"`
	PRODUCTINGREDIENTS       string `json:"productINGREDIENTS"`
	PRODUCTCATEGORY          string `json:"productCATEGORY"`
	PRODUCTIMAGE1            string `json:"productIMAGE1"`
	PRODUCTIMAGE2            string `json:"productIMAGE2"`
}

type productListPageEdit struct {
	PRODUCTID                int    `json:"productID"`
	PRODUCTNAME              string `json:"productNAME"`
	PRODUCTGUID              string `json:"productGUID"`
	PRODUCTDESCRIPTION       string `json:"productDESCRIPTION"`
	PRODUCTMANUFACTURER      string `json:"productMANUFACTURER"`
	PRODUCTMANUFACTURERADDR  string `json:"productMANUFACTURERADDR"`
	PRODUCTLIKES             string `json:"productLIKES"`
	PRODUCTDISLIKES          string `json:"productDISLIKES"`
	PRODUCTTREND             string `json:"productTREND"`
	PRODUCTTRENDDIRECTION    string `json:"productTRENDDIRECTION"`
	PRODUCTSENTIMENT         string `json:"productSENTIMENT"`
	PRODUCTSENTIMENTMOOD     string `json:"productSENTIMENTMOOD"`
	PRODUCTUSERCOMMENTS      string `json:"productUSERCOMMENTS"`
	PRODUCTRATING            string `json:"productRATING"`
	PRODUCTAUTHOR            string `json:"productAUTHOR"`
	PRODUCTDATEPUBLISHED     string `json:"productDATEPUBLISHED"`
	PRODUCTSKU               string `json:"productSKU"`
	PRODUCTMPN               string `json:"productMPN"`
	PRODUCTPRICE             string `json:"productPRICE"`
	PRODUCTLOCATIONCOUNT     string `json:"productLOCATIONCOUNT"`
	PRODUCTFIRSTCOMPETITION  string `json:"productFIRSTCOMPETITION"`
	PRODUCTSECONDCOMPETITION string `json:"productSECONDCOMPETITION"`
	PRODUCTINGREDIENTS       string `json:"productINGREDIENTS"`
	PRODUCTCATEGORY          string `json:"productCATEGORY"`
	PRODUCTIMAGE1            string `json:"productIMAGE1"`
	PRODUCTIMAGE2            string `json:"productIMAGE2"`
}

type productRichSnippet struct {
	PRODUCTID           int    `json:"productID"`
	PRODUCTNAME         string `json:"productNAME"`
	PRODUCTGUID         string `json:"productGUID"`
	PRODUCTDESCRIPTION  string `json:"productDESCRIPTION"`
	PRODUCTMANUFACTURER string `json:"productMANUFACTURER"`
	PRODUCTUSERCOMMENTS string `json:"productUSERCOMMENTS"`
	PRODUCTRATING       string `json:"productRATING"`
	PRODUCTAUTHOR       string `json:"productAUTHOR"`
	PRODUCTFULLCOMMENTS string `json:"productCOMMENTS"`
	PRODUCTSKU          string `json:"productSKU"`
	PRODUCTMPN          string `json:"productMPN"`
	PRODUCTPRICE        string `json:"productPRICE"`
	PRODUCTIMAGE1       string `json:"productIMAGE1"`
}

type commentList struct {
	PRODUCTNAME   string `json:"productname"`
	LIKES         string `json:"likes"`
	DISLIKE       string `json:"dislike"`
	PRODUCTRATING string `json:"rate"`
	USERCOMMENTS  string `json:"usercomments"`
	LATITUDE      string `json:"latitude"`
	LONGITUDE     string `json:"longitude"`
	DATEPUBLISHED string `json:"datePublished"`
	AUTHOR        string `json:"author"`
}

type category struct {
	ID       int    `json:"id"`
	CATEGORY string `json:"category"`
}

type competitorsI struct {
	PRODUCTID   int    `json:"id"`
	PRODUCTGUID string `json:"productguid"`
	PRODUCTNAME string `json:"productname"`
}

type competitors struct {
	PRODUCTID             int    `json:"productID"`
	PRODUCTNAME           string `json:"productNAME"`
	PRODUCTGUID           string `json:"productGUID"`
	PRODUCTDESCRIPTION    string `json:"productDESCRIPTION"`
	PRODUCTMANUFACTURER   string `json:"productMANUFACTURER"`
	PRODUCTLIKES          string `json:"productLIKES"`
	PRODUCTDISLIKES       string `json:"productDISLIKES"`
	PRODUCTTREND          string `json:"productTREND"`
	PRODUCTTRENDDIRECTION string `json:"productTRENDDIRECTION"`
	PRODUCTSENTIMENT      string `json:"productSENTIMENT"`
	PRODUCTSENTIMENTMOOD  string `json:"productSENTIMENTMOOD"`
	PRODUCTUSERCOMMENTS   string `json:"productUSERCOMMENTS"`
	PRODUCTRATING         string `json:"productRATING"`
	PRODUCTDATEPUBLISHED  string `json:"productDATEPUBLISHED"`
	PRODUCTPRICE          string `json:"productPRICE"`
	PRODUCTLOCATIONCOUNT  string `json:"productLOCATIONCOUNT"`
	PRODUCTINGREDIENTS    string `json:"productINGREDIENTS"`
	PRODUCTCATEGORY       string `json:"productCATEGORY"`
	PRODUCTIMAGE1         string `json:"productIMAGE1"`
	PRODUCTIMAGE2         string `json:"productIMAGE2"`
}

type cU struct {
	COUNT int `json:"count"`
}

func getAllItemsFrontPage() ([]productList, error) {
	var productLists = []productList{}
	var (
		singleItem productList
	)
	row, err := database.DB.Query(`
		SELECT all_products.id AS product_ID, 
		all_products.title AS productname,
		all_products.product_name_clean_url AS productGUID, 
		all_products.about AS description, 
		all_products.manufacturer AS manufacturer, 
		SUM(product_review.likes) AS likes, 
		SUM(product_review.dislikes) AS dislikes, 
		CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'trending_up' ELSE 'trending_down' END AS trend, 
		CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'Up' ELSE 'Down' END AS trend_direction, 
		CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'sentiment_very_satisfied' ELSE 'sentiment_very_dissatisfied' END AS sentiment, 
		CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'Good' ELSE 'Bad' END AS sentiment_mood, 
		COUNT(product_review.user_comments) AS usercomments, 
		AVG(product_review.rating) AS rating,
		product_review.date AS datePublished, 
		all_products.price AS price, 
		COUNT(user_location_lat) + COUNT(user_location_lon) AS locationcount, 
		all_products.ingredients AS ingredients, 
		product_categories.category AS category, 
		CONCAT('https://asknigeria.com.ng/assets/brands/images/281x224/', SUBSTR(all_products.product_image_1,8)) AS productImage_1, 
		CONCAT('https://asknigeria.com.ng/assets/brands/images/750x224/', SUBSTR(all_products.product_image_2,8)) AS productImage_2 
		FROM all_products 
		JOIN product_review ON all_products.id = product_review.product_id
		JOIN product_categories ON all_products.category = product_categories.id  
		WHERE
		all_products.about IS NOT NULL AND
		all_products.manufacturer IS NOT NULL AND 
		all_products.address IS NOT NULL AND
		all_products.ingredients IS NOT NULL AND
		all_products.product_image_1 IS NOT NULL AND
		all_products.product_image_2 IS NOT NULL AND
		all_products.price IS NOT NULL
		GROUP BY all_products.id ORDER BY rating DESC
	`)
	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&singleItem.PRODUCTID,
				&singleItem.PRODUCTNAME,
				&singleItem.PRODUCTGUID,
				&singleItem.PRODUCTDESCRIPTION,
				&singleItem.PRODUCTMANUFACTURER,
				&singleItem.PRODUCTLIKES,
				&singleItem.PRODUCTDISLIKES,
				&singleItem.PRODUCTTREND,
				&singleItem.PRODUCTTRENDDIRECTION,
				&singleItem.PRODUCTSENTIMENT,
				&singleItem.PRODUCTSENTIMENTMOOD,
				&singleItem.PRODUCTUSERCOMMENTS,
				&singleItem.PRODUCTRATING,
				&singleItem.PRODUCTDATEPUBLISHED,
				&singleItem.PRODUCTPRICE,
				&singleItem.PRODUCTLOCATIONCOUNT,
				&singleItem.PRODUCTINGREDIENTS,
				&singleItem.PRODUCTCATEGORY,
				&singleItem.PRODUCTIMAGE1,
				&singleItem.PRODUCTIMAGE2,
			)
			if err != nil {
				return nil, err
			}
			productLists = append(productLists, singleItem)
		}
		defer row.Close()
	}
	return productLists, nil
}

// Return a list of all the products for ld/json
func getAllItems() ([]itemList, string) {
	var itemLists = []itemList{}
	var (
		singularItem itemList
	)
	row, err := database.DB.Query(`
		SELECT all_products.id AS product_ID, 
		all_products.product_name_clean_url AS productGUID 
		FROM all_products
		JOIN product_review ON all_products.id = product_review.product_id
		JOIN product_categories ON all_products.category = product_categories.id
		WHERE
		all_products.about IS NOT NULL AND
		all_products.manufacturer IS NOT NULL AND 
		all_products.address IS NOT NULL AND
		all_products.ingredients IS NOT NULL AND
		all_products.product_image_1 IS NOT NULL AND
		all_products.product_image_2 IS NOT NULL AND
		all_products.price IS NOT NULL
		GROUP BY all_products.id ORDER BY rating DESC
	`)
	if err != nil {
		return nil, err.Error()
	} else {
		for row.Next() {
			err = row.Scan(&singularItem.PRODUCTID, &singularItem.PRODUCTGUID)
			if err != nil {
				return nil, err.Error()
			}
			itemLists = append(itemLists, singularItem)
		}
		defer row.Close()
	}
	return itemLists, ""
}

func getProductLists() ([]itemProductList, error) {
	var itemProductLists = []itemProductList{}
	var (
		singularItem itemProductList
	)
	row, err := database.DB.Query(`
		SELECT all_products.id AS product_ID, 
		all_products.title AS productTITLE ,
		all_products.product_name_clean_url AS productGUID ,
		all_products.manufacturer AS manufacturer,
		product_categories.category AS category
		FROM all_products
		JOIN product_review ON all_products.id = product_review.product_id
		JOIN product_categories ON all_products.category = product_categories.id
		WHERE
		all_products.about IS NOT NULL AND
		all_products.manufacturer IS NOT NULL AND 
		all_products.address IS NOT NULL AND
		all_products.ingredients IS NOT NULL AND
		all_products.product_image_1 IS NOT NULL AND
		all_products.product_image_2 IS NOT NULL AND
		all_products.price IS NOT NULL
		GROUP BY all_products.id ORDER BY rating DESC
	`)
	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(&singularItem.PRODUCTID, &singularItem.PRODUCTTITLE, &singularItem.PRODUCTGUID,
				&singularItem.PRODUCTMANUFACTURER, &singularItem.PRODUCTCATEGORY)
			if err != nil {
				return nil, err
			}
			itemProductLists = append(itemProductLists, singularItem)
		}
		defer row.Close()
	}
	return itemProductLists, nil
}

func getProductListsByUser(userid string) ([]userRatedProductList, error) {
	var itemProductLists = []userRatedProductList{}
	var (
		singularItem userRatedProductList
	)
	row, err := database.DB.Query(`
		SELECT all_products.id AS product_ID, 
		product_review.id AS reviewID,
		all_products.title AS productTITLE ,
		product_review.rating AS rating,
		all_products.product_name_clean_url AS productGUID ,
		all_products.manufacturer AS manufacturer,
		product_categories.category AS category
		FROM all_products
		JOIN product_review ON all_products.id = product_review.product_id
		JOIN product_categories ON all_products.category = product_categories.id
		WHERE
		product_review.user = ? AND
		all_products.about IS NOT NULL AND
		all_products.manufacturer IS NOT NULL AND 
		all_products.address IS NOT NULL AND
		all_products.ingredients IS NOT NULL AND
		all_products.product_image_1 IS NOT NULL AND
		all_products.product_image_2 IS NOT NULL AND
		all_products.price IS NOT NULL
		GROUP BY all_products.id ORDER BY rating DESC
	`, userid)
	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(&singularItem.PRODUCTID,
				&singularItem.REVIEWID,
				&singularItem.PRODUCTTITLE,
				&singularItem.RATING,
				&singularItem.PRODUCTGUID,
				&singularItem.PRODUCTMANUFACTURER, &singularItem.PRODUCTCATEGORY)
			if err != nil {
				return nil, err
			}
			itemProductLists = append(itemProductLists, singularItem)
		}
		defer row.Close()
	}
	return itemProductLists, nil
}

func getAllItemsCount() int {
	var (
		numCount int
	)
	row, err := database.DB.Query(`
		SELECT COUNT(*) FROM 
			(SELECT COUNT(*) AS count 
			FROM all_products 
			JOIN product_review ON all_products.id = product_review.product_id 
			JOIN product_categories ON all_products.category = product_categories.id 
			WHERE all_products.about IS NOT NULL 
			AND all_products.manufacturer IS NOT NULL 
			AND  all_products.address IS NOT NULL 
			AND all_products.ingredients IS NOT NULL 
			AND all_products.product_image_1 IS NOT NULL 
			AND all_products.product_image_2 IS NOT NULL 
			AND all_products.price IS NOT NULL 
			GROUP BY all_products.id ORDER BY rating DESC) 
		AS count
	`)
	numCount = checkCount(row)
	checkErr(err)
	return numCount
}

func getAllCommentsCount(productID int) int {
	var (
		numCount int
	)
	row, err := database.DB.Query(`
		SELECT COUNT(*) 
		FROM 
		product_review WHERE product_id = ?
	`, productID)
	numCount = checkCount(row)
	checkErr(err)
	return numCount
}

func getPID(productID string) (int, error) {
	var (
		pid int
	)
	row, err := database.DB.Query(`
		SELECT id FROM all_products 
		WHERE product_name_clean_url = ?
	`, productID)
	if err != nil {
		return 0, err
	}
	for row.Next() {
		err := row.Scan(&pid)
		if err != nil {
			return 0, err
		}
	}
	return pid, nil
}

func getAllComments(productID int) ([]commentList, error) {
	var commentLists = []commentList{}
	var (
		comment commentList
	)

	row, err := database.DB.Query(`
			SELECT all_products.title AS productname, 
			product_review.likes AS likes, 
			product_review.dislikes AS dislike, 
			product_review.rating AS rate, 
			product_review.user_comments AS comment, 
			product_review.user_location_lat AS latitude, 
			product_review.user_location_lon AS longitude, 
			product_review.date AS datePublished, 
			product_review.author AS author 
			FROM product_review 
			JOIN all_products 
			ON product_review.product_id = all_products.ID 
			WHERE product_review.product_id = ?
	`, productID)
	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&comment.PRODUCTNAME,
				&comment.LIKES,
				&comment.DISLIKE,
				&comment.PRODUCTRATING,
				&comment.USERCOMMENTS,
				&comment.LATITUDE,
				&comment.LONGITUDE,
				&comment.DATEPUBLISHED,
				&comment.AUTHOR,
			)
			if err != nil {
				return nil, err
			}
			commentLists = append(commentLists, comment)
		}
		defer row.Close()
	}
	return commentLists, nil
}

func getCommentsData(pid string) {

}

func getProductData(pid string) ([]productListPage, error) {
	var productPage = []productListPage{}
	var (
		singleItem productListPage
	)
	row, err := database.DB.Query(`
		SELECT all_products.id AS product_ID, 
		all_products.title AS productname,
		all_products.product_name_clean_url AS productGUID, 
		all_products.about AS description, 
		all_products.manufacturer AS manufacturer, 
		SUM(product_review.likes) AS likes, 
		SUM(product_review.dislikes) AS dislikes, 
		CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'trending_up' ELSE 'trending_down' END AS trend, 
		CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'Up' ELSE 'Down' END AS trend_direction, 
		CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'sentiment_very_satisfied' ELSE 'sentiment_very_dissatisfied' END AS sentiment, 
		CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'Good' ELSE 'Bad' END AS sentiment_mood, 
		COUNT(product_review.user_comments) AS usercomments, 
		AVG(product_review.rating) AS rating,
		product_review.author AS author,
		UNIX_TIMESTAMP(product_review.date) AS datePublished, 
		all_products.sku AS sku, 
		all_products.mpn AS mpn,
		all_products.price AS price, 
		COUNT(user_location_lat) + COUNT(user_location_lon) AS locationcount, 
		all_products.competitor_1 AS firstCompetition, 
		all_products.competitor_2 AS secondCompetition,
		all_products.ingredients AS ingredients, 
		product_categories.category AS category, 
		CONCAT('https://asknigeria.com.ng/assets/brands/images/281x224/', SUBSTR(all_products.product_image_1,8)) AS productImage_1, 
		CONCAT('https://asknigeria.com.ng/assets/brands/images/750x224/', SUBSTR(all_products.product_image_2,8)) AS productImage_2 
		FROM all_products 
		JOIN product_review ON all_products.id = product_review.product_id
		JOIN product_categories ON all_products.category = product_categories.id  
		WHERE 
		all_products.about <> '' 
		AND all_products.about IS NOT NULL
		AND all_products.product_name_clean_url = ?
		AND all_products.manufacturer IS NOT NULL AND 
		all_products.address IS NOT NULL AND
		all_products.ingredients IS NOT NULL AND
		all_products.product_image_1 IS NOT NULL AND
		all_products.product_image_2 IS NOT NULL AND
		all_products.price IS NOT NULL
		GROUP BY all_products.id ORDER BY rating DESC
	`, pid)
	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&singleItem.PRODUCTID,
				&singleItem.PRODUCTNAME,
				&singleItem.PRODUCTGUID,
				&singleItem.PRODUCTDESCRIPTION,
				&singleItem.PRODUCTMANUFACTURER,
				&singleItem.PRODUCTLIKES,
				&singleItem.PRODUCTDISLIKES,
				&singleItem.PRODUCTTREND,
				&singleItem.PRODUCTTRENDDIRECTION,
				&singleItem.PRODUCTSENTIMENT,
				&singleItem.PRODUCTSENTIMENTMOOD,
				&singleItem.PRODUCTUSERCOMMENTS,
				&singleItem.PRODUCTRATING,
				&singleItem.PRODUCTAUTHOR,
				&singleItem.PRODUCTDATEPUBLISHED,
				&singleItem.PRODUCTSKU,
				&singleItem.PRODUCTMPN,
				&singleItem.PRODUCTPRICE,
				&singleItem.PRODUCTLOCATIONCOUNT,
				&singleItem.PRODUCTFIRSTCOMPETITION,
				&singleItem.PRODUCTSECONDCOMPETITION,
				&singleItem.PRODUCTINGREDIENTS,
				&singleItem.PRODUCTCATEGORY,
				&singleItem.PRODUCTIMAGE1,
				&singleItem.PRODUCTIMAGE2,
			)
			if err != nil {
				return nil, err
			}
			productPage = append(productPage, singleItem)
		}
		defer row.Close()
	}
	return productPage, nil
}

func getRichSnippet(pid string) ([]productRichSnippet, error) {

	var richSnippet = []productRichSnippet{}
	var (
		singleItem productRichSnippet
	)

	row, err := database.DB.Query(`
		SELECT all_products.id AS product_ID, 
		all_products.title AS productname,
		all_products.product_name_clean_url AS productGUID, 
		all_products.about AS description, 
		all_products.manufacturer AS manufacturer, 
		COUNT(product_review.user_comments) AS usercomments, 
		AVG(product_review.rating) AS rating,
		product_review.author AS author,
		GROUP_CONCAT(
			CONCAT_WS(
					'#', 
					product_review.likes, 
					product_review.dislikes, 
					product_review.rating, 
					product_review.user_comments, 
					product_review.author, 
					product_review.user_location_lat, 
					product_review.user_location_lon, 
					UNIX_TIMESTAMP(product_review.date), 
					'~'
				)
		) AS fullcomments,
		all_products.sku AS sku, 
		all_products.mpn AS mpn,
		all_products.price AS price, 
		CONCAT('https://asknigeria.com.ng/assets/brands/images/281x224/', SUBSTR(all_products.product_image_1,8)) AS productImage_1
		FROM all_products 
		JOIN product_review ON all_products.id = product_review.product_id
		JOIN product_categories ON all_products.category = product_categories.id  
		WHERE 
		all_products.product_name_clean_url = ? AND
		all_products.about <> '' 
		AND all_products.about IS NOT NULL
		AND all_products.manufacturer IS NOT NULL AND 
		all_products.address IS NOT NULL AND
		all_products.ingredients IS NOT NULL AND
		all_products.product_image_1 IS NOT NULL AND
		all_products.product_image_2 IS NOT NULL AND
		all_products.price IS NOT NULL
		GROUP BY all_products.id ORDER BY rating DESC
	`, pid)
	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&singleItem.PRODUCTID,
				&singleItem.PRODUCTNAME,
				&singleItem.PRODUCTGUID,
				&singleItem.PRODUCTDESCRIPTION,
				&singleItem.PRODUCTMANUFACTURER,
				&singleItem.PRODUCTUSERCOMMENTS,
				&singleItem.PRODUCTRATING,
				&singleItem.PRODUCTAUTHOR,
				&singleItem.PRODUCTFULLCOMMENTS,
				&singleItem.PRODUCTSKU,
				&singleItem.PRODUCTMPN,
				&singleItem.PRODUCTPRICE,
				&singleItem.PRODUCTIMAGE1,
			)
			if err != nil {
				return nil, err
			}
			richSnippet = append(richSnippet, singleItem)
		}
		defer row.Close()
	}

	return richSnippet, nil
}

func getCategories() ([]category, error) {
	var categoryLists = []category{}
	var (
		c category
	)

	row, err := database.DB.Query(`
			SELECT id, category 
			FROM product_categories`)
	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&c.ID,
				&c.CATEGORY,
			)
			if err != nil {
				return nil, err
			}
			categoryLists = append(categoryLists, c)
		}
		defer row.Close()
	}
	return categoryLists, nil
}

func makeProduct(items Myform) (*Myform, error) {
	// We need to remember to generate sku and mpn
	stmt, err := database.DB.Prepare(`insert into all_products
				(
					product_name_clean_url, title, category, competitor_1, competitor_2,
					about, manufacturer, address, ingredients, updated, product_image_1,
					product_image_2, price, sku, mpn
				)
				values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	res, err := stmt.Exec(
		items.PRODUCTGUID, items.PRODUCTNAME, items.CATEGORIES, items.COMPETITORS[0],
		items.COMPETITORS[1], items.ABOUT, items.MANUFACTURER, items.MANUFACTURERADDRESS,
		items.INGREDIENTS, 0, items.PRODUCTIMAGE1, items.PRODUCTIMAGE2, items.PRICE,
		items.SKU, items.MPN,
	)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	lid, err := res.LastInsertId()

	var datetime = time.Now()
	datetime.Format(time.RFC3339)
	stmtX, errX := database.DB.Prepare(`insert into product_review
				(
					product_id, product_category, likes, dislikes, rating, user_comments, 
					user_location_lat, user_location_lon, date, author
				)
				values(?,?,?,?,?,?,?,?,?,?);`)
	if errX != nil {
		return nil, errors.New(errX.Error())
	}
	_, errX = stmtX.Exec(
		lid, items.CATEGORIES, 0, 0, 1, "Awesome Product", "", "", datetime, "System Generated",
	)

	if errX != nil {
		return nil, errors.New(errX.Error())
	}
	defer stmt.Close()
	defer stmtX.Close()

	return &items, nil
}

func genSKU() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	v := r.Int()
	return v
}

func genMPN() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	v := r.Int()
	return v
}

func getCompetitorsI() ([]competitorsI, error) {
	var competitorsListsI = []competitorsI{}
	var (
		c competitorsI
	)

	row, err := database.DB.Query(`
			SELECT id, product_name_clean_url, title 
			FROM all_products`)
	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&c.PRODUCTID,
				&c.PRODUCTGUID,
				&c.PRODUCTNAME,
			)
			if err != nil {
				return nil, err
			}
			competitorsListsI = append(competitorsListsI, c)
		}
		defer row.Close()
	}
	return competitorsListsI, nil
}

func getCompetitors(guid string) ([]competitors, []noCompetition, error) {
	guid = genGUID(guid)
	var competitorsLists = []competitors{}
	var noCompetitionListsQ = []noCompetition{}
	var (
		singleItem competitors
		z          noCompetition
	)

	if errX := database.DB.QueryRow(`
					SELECT all_products.id AS product_ID, 
					all_products.title AS productname,
					all_products.product_name_clean_url AS productGUID, 
					all_products.about AS description, 
					all_products.manufacturer AS manufacturer, 
					SUM(product_review.likes) AS likes, 
					SUM(product_review.dislikes) AS dislikes, 
					CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'trending_up' ELSE 'trending_down' END AS trend, 
					CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'Up' ELSE 'Down' END AS trend_direction, 
					CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'sentiment_very_satisfied' ELSE 'sentiment_very_dissatisfied' END AS sentiment, 
					CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'Good' ELSE 'Bad' END AS sentiment_mood, 
					COUNT(product_review.user_comments) AS usercomments, 
					AVG(product_review.rating) AS rating,
					UNIX_TIMESTAMP(product_review.date) AS datePublished, 
					all_products.price AS price, 
					COUNT(user_location_lat) + COUNT(user_location_lon) AS locationcount, 
					all_products.ingredients AS ingredients, 
					product_categories.category AS category, 
					CONCAT('https://asknigeria.com.ng/assets/brands/images/281x224/', SUBSTR(all_products.product_image_1,8)) AS productImage_1, 
					CONCAT('https://asknigeria.com.ng/assets/brands/images/750x224/', SUBSTR(all_products.product_image_2,8)) AS productImage_2 
					FROM all_products 
					JOIN product_review ON all_products.id = product_review.product_id
					JOIN product_categories ON all_products.category = product_categories.id  
					WHERE 
					all_products.about <> '' 
					AND all_products.about IS NOT NULL
					AND all_products.product_name_clean_url = ?
					AND all_products.manufacturer IS NOT NULL AND 
					all_products.address IS NOT NULL AND
					all_products.ingredients IS NOT NULL AND
					all_products.product_image_1 IS NOT NULL AND
					all_products.product_image_2 IS NOT NULL AND
					all_products.price IS NOT NULL
					GROUP BY all_products.id ORDER BY rating DESC
				`, guid).Scan(
		&singleItem.PRODUCTID,
		&singleItem.PRODUCTNAME,
		&singleItem.PRODUCTGUID,
		&singleItem.PRODUCTDESCRIPTION,
		&singleItem.PRODUCTMANUFACTURER,
		&singleItem.PRODUCTLIKES,
		&singleItem.PRODUCTDISLIKES,
		&singleItem.PRODUCTTREND,
		&singleItem.PRODUCTTRENDDIRECTION,
		&singleItem.PRODUCTSENTIMENT,
		&singleItem.PRODUCTSENTIMENTMOOD,
		&singleItem.PRODUCTUSERCOMMENTS,
		&singleItem.PRODUCTRATING,
		&singleItem.PRODUCTDATEPUBLISHED,
		&singleItem.PRODUCTPRICE,
		&singleItem.PRODUCTLOCATIONCOUNT,
		&singleItem.PRODUCTINGREDIENTS,
		&singleItem.PRODUCTCATEGORY,
		&singleItem.PRODUCTIMAGE1,
		&singleItem.PRODUCTIMAGE2,
	); errX == nil {
		// Meaning we have one result
		competitorsLists = append(competitorsLists, singleItem)
	} else if errX == sql.ErrNoRows {
		// Meaning no result
		errZ := database.DB.QueryRow(`
					SELECT
					all_products.title AS productname, 
					all_products.product_name_clean_url AS productGUID, 
					product_categories.category AS category FROM all_products 
					JOIN product_categories ON all_products.category = product_categories.id 
					 WHERE all_products.product_name_clean_url = ?
				`, guid).Scan(&z.PRODUCTNAME, &z.PRODUCTGUID, &z.PRODUCTCATEGORY)
		if errZ != nil {
			return nil, nil, errZ
		}
		noCompetitionListsQ = append(noCompetitionListsQ, z)

	} else {
		return nil, nil, errX
	}

	return competitorsLists, noCompetitionListsQ, nil
}

func canComment(pid, cat, username string) bool {
	var (
		numCount int
		b        cU
	)

	sqlRaw := fmt.Sprintf(`SELECT COUNT(*) AS count 
			FROM product_review 
			WHERE product_id = '%s'
			AND product_category = '%s'
			AND user = '%s'`, pid, cat, username)

	if errX := database.DB.QueryRow(sqlRaw).Scan(&b.COUNT); errX == nil {
		numCount = b.COUNT
	} else if errX == sql.ErrNoRows {
		log.Println(sql.ErrNoRows)
	} else {
		log.Println(errX)
	}
	if numCount > 0 {
		return false
	} else {
		return true
	}
}

func insertComments(pid, cat, username, comments,
	rating, sentiment, latitude, longitude, author string) (int, error) {
	if canComment(pid, cat, username) {
		var datetime = time.Now()
		datetime.Format(time.RFC3339)
		var like int
		var dislike int
		if sentiment == "like" {
			like = 1
		} else {
			dislike = 1
		}
		stmtX, errX := database.DB.Prepare(`insert into product_review
				(
					product_id, product_category, likes, dislikes, rating, 
					user_comments, 
					user_location_lat, user_location_lon, date, user, author
				)
				values(?,?,?,?,?,?,?,?,?,?,?);`)
		if errX != nil {
			return 0, errors.New(errX.Error())
		}
		resX, errX := stmtX.Exec(
			pid, cat, like, dislike, rating, comments, latitude, longitude, datetime, username, author,
		)

		if errX != nil {
			return 0, errors.New(errX.Error())
		}

		defer stmtX.Close()

		lid, errX := resX.LastInsertId()

		if errX != nil {
			return 0, errors.New(errX.Error())
		}

		return int(lid), nil
	} else {
		return 0, errors.New("Oops, you cannot comment twice on the same product")
	}

}

type productLikes struct {
	PRODUCTNAME string `json:"productname"`
	LIKES       string `json:"likes"`
}

type productDisLikes struct {
	PRODUCTNAME string `json:"productname"`
	DISLIKES    string `json:"dislikes"`
}

type reviewRating struct {
	PRODUCTNAME string `json:"productname"`
	RATING      string `json:"rating"`
}

type acceptanceArea struct {
	LATITUDE  float32 `json:"latitude"`
	LONGITUDE float32 `json:"longitude"`
	RATING    int     `json:"rating"`
}

type rejectionArea struct {
	LATITUDE  float32 `json:"latitude"`
	LONGITUDE float32 `json:"longitude"`
	RATING    int     `json:"rating"`
}

func getReviewLikes(data []string) ([]productLikes, error) {
	var likesList = []productLikes{}
	var (
		j productLikes
	)

	ids := strings.Join(data, "','")

	sqlRaw := fmt.Sprintf(`SELECT all_products.title AS productname, 
				IFNULL(SUM(product_review.likes),0) AS likes
				FROM all_products 
				JOIN product_review 
				ON all_products.id = product_review.product_id  
				WHERE all_products.title IN ('%s')  GROUP BY all_products.title`, ids)

	row, err := database.DB.Query(sqlRaw)

	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&j.PRODUCTNAME,
				&j.LIKES,
			)
			if err != nil {
				return nil, err
			}
			likesList = append(likesList, j)
		}
		defer row.Close()
	}
	return likesList, nil
}

func getReviewDisLikes(data []string) ([]productDisLikes, error) {
	var dislikesList = []productDisLikes{}
	var (
		v productDisLikes
	)

	ids := strings.Join(data, "','")

	sqlRaw := fmt.Sprintf(`SELECT all_products.title AS productname, 
				IFNULL(SUM(product_review.dislikes),0) AS dislikes
				FROM all_products 
				JOIN product_review 
				ON all_products.id = product_review.product_id  
				WHERE all_products.title IN ('%s')  GROUP BY all_products.title`, ids)

	row, err := database.DB.Query(sqlRaw)

	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&v.PRODUCTNAME,
				&v.DISLIKES,
			)
			if err != nil {
				return nil, err
			}
			dislikesList = append(dislikesList, v)
		}
		defer row.Close()
	}
	return dislikesList, nil
}

func getReviewRating(data []string) ([]reviewRating, error) {
	var ratingsList = []reviewRating{}
	var (
		z reviewRating
	)

	ids := strings.Join(data, "','")

	sqlRaw := fmt.Sprintf(`SELECT all_products.title AS productname, 
				IFNULL(SUM(product_review.rating),0) AS rating
				FROM all_products 
				JOIN product_review 
				ON all_products.id = product_review.product_id  
				WHERE all_products.title IN ('%s')  GROUP BY all_products.title`, ids)

	row, err := database.DB.Query(sqlRaw)

	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&z.PRODUCTNAME,
				&z.RATING,
			)
			if err != nil {
				return nil, err
			}
			ratingsList = append(ratingsList, z)
		}
		defer row.Close()
	}
	return ratingsList, nil
}

func getAcceptanceAreas(pid string) ([]acceptanceArea, error) {
	var acceptanceList = []acceptanceArea{}
	var (
		z acceptanceArea
	)
	row, err := database.DB.Query(`
		SELECT user_location_lat AS latitude, user_location_lon AS longitude,
		rating AS rate FROM product_review WHERE product_id = ?
		 AND rating BETWEEN '2.5' AND '5' 
	`, pid)

	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&z.LATITUDE,
				&z.LONGITUDE,
				&z.RATING,
			)
			if err != nil {
				return nil, err
			}
			acceptanceList = append(acceptanceList, z)
		}
		defer row.Close()
	}
	return acceptanceList, nil
}

func getRejectionAreas(pid string) ([]rejectionArea, error) {
	var rejectionList = []rejectionArea{}
	var (
		z rejectionArea
	)
	row, err := database.DB.Query(`
			SELECT user_location_lat AS latitude, user_location_lon AS longitude,
			rating AS rate FROM product_review WHERE product_id = ?
			AND rating BETWEEN '0' AND '2.5' 
	`, pid)

	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&z.LATITUDE,
				&z.LONGITUDE,
				&z.RATING,
			)
			if err != nil {
				return nil, err
			}
			rejectionList = append(rejectionList, z)
		}
		defer row.Close()
	}
	return rejectionList, nil
}

type pRecommendation struct {
	PRODUCTID             int    `json:"productID"`
	PRODUCTNAME           string `json:"productNAME"`
	PRODUCTGUID           string `json:"productGUID"`
	PRODUCTDESCRIPTION    string `json:"productDESCRIPTION"`
	PRODUCTMANUFACTURER   string `json:"productMANUFACTURER"`
	PRODUCTLIKES          string `json:"productLIKES"`
	PRODUCTDISLIKES       string `json:"productDISLIKES"`
	PRODUCTTREND          string `json:"productTREND"`
	PRODUCTTRENDDIRECTION string `json:"productTRENDDIRECTION"`
	PRODUCTSENTIMENT      string `json:"productSENTIMENT"`
	PRODUCTSENTIMENTMOOD  string `json:"productSENTIMENTMOOD"`
	PRODUCTUSERCOMMENTS   string `json:"productUSERCOMMENTS"`
	PRODUCTRATING         string `json:"productRATING"`
	PRODUCTDATEPUBLISHED  string `json:"productDATEPUBLISHED"`
	PRODUCTPRICE          string `json:"productPRICE"`
	PRODUCTLOCATIONCOUNT  string `json:"productLOCATIONCOUNT"`
	PRODUCTINGREDIENTS    string `json:"productINGREDIENTS"`
	PRODUCTCATEGORY       string `json:"productCATEGORY"`
	PRODUCTIMAGE1         string `json:"productIMAGE1"`
	PRODUCTIMAGE2         string `json:"productIMAGE2"`
}

type tMp struct {
	PRODUCTNAME string `json:"productname"`
	PRODUCTGUID string `json:"productguid"`
	RATING      string `json:"rating"`
	LIKES       string `json:"likes"`
}

type noCompetition struct {
	PRODUCTNAME     string `json:"productname"`
	PRODUCTGUID     string `json:"productguid"`
	PRODUCTCATEGORY string `json:"category"`
}

func productRecommendation(data []string) ([]pRecommendation, []noCompetition, error) {
	var pRecommendationLists = []pRecommendation{}
	var noCompetitionLists = []noCompetition{}
	var (
		singleItem pRecommendation
		j          tMp
		x          noCompetition
	)

	ids := strings.Join(data, "','")

	sqlRaw := fmt.Sprintf(`SELECT all_products.title AS productname, 
				all_products.product_name_clean_url AS productGUID, 
				AVG(product_review.rating) AS rating, AVG(product_review.likes) AS likes 
				FROM all_products 
				JOIN product_review ON all_products.id = product_review.product_id 
				WHERE all_products.product_name_clean_url IN ('%s')  
				ORDER BY product_review.rating DESC LIMIT 1`, ids)

	row, err := database.DB.Query(sqlRaw)

	if err != nil {
		return nil, nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&j.PRODUCTNAME,
				&j.PRODUCTGUID,
				&j.RATING,
				&j.LIKES,
			)
			if err != nil {
				return nil, nil, err
			}

			// Second Query
			/////////////////////////////////////////////////////////////////////

			if errX := database.DB.QueryRow(`
					SELECT all_products.id AS product_ID, 
					all_products.title AS productname,
					all_products.product_name_clean_url AS productGUID, 
					all_products.about AS description, 
					all_products.manufacturer AS manufacturer, 
					SUM(product_review.likes) AS likes, 
					SUM(product_review.dislikes) AS dislikes, 
					CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'trending_up' ELSE 'trending_down' END AS trend, 
					CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'Up' ELSE 'Down' END AS trend_direction, 
					CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'sentiment_very_satisfied' ELSE 'sentiment_very_dissatisfied' END AS sentiment, 
					CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'Good' ELSE 'Bad' END AS sentiment_mood, 
					COUNT(product_review.user_comments) AS usercomments, 
					AVG(product_review.rating) AS rating,
					UNIX_TIMESTAMP(product_review.date) AS datePublished, 
					all_products.price AS price, 
					COUNT(user_location_lat) + COUNT(user_location_lon) AS locationcount, 
					all_products.ingredients AS ingredients, 
					product_categories.category AS category, 
					CONCAT('https://asknigeria.com.ng/assets/brands/images/281x224/', SUBSTR(all_products.product_image_1,8)) AS productImage_1, 
					CONCAT('https://asknigeria.com.ng/assets/brands/images/750x224/', SUBSTR(all_products.product_image_2,8)) AS productImage_2 
					FROM all_products 
					JOIN product_review ON all_products.id = product_review.product_id
					JOIN product_categories ON all_products.category = product_categories.id  
					WHERE 
					all_products.about <> '' 
					AND all_products.about IS NOT NULL
					AND all_products.product_name_clean_url = ?
					AND all_products.manufacturer IS NOT NULL AND 
					all_products.address IS NOT NULL AND
					all_products.ingredients IS NOT NULL AND
					all_products.product_image_1 IS NOT NULL AND
					all_products.product_image_2 IS NOT NULL AND
					all_products.price IS NOT NULL
					GROUP BY all_products.id ORDER BY rating DESC
				`, &j.PRODUCTGUID).Scan(
				&singleItem.PRODUCTID,
				&singleItem.PRODUCTNAME,
				&singleItem.PRODUCTGUID,
				&singleItem.PRODUCTDESCRIPTION,
				&singleItem.PRODUCTMANUFACTURER,
				&singleItem.PRODUCTLIKES,
				&singleItem.PRODUCTDISLIKES,
				&singleItem.PRODUCTTREND,
				&singleItem.PRODUCTTRENDDIRECTION,
				&singleItem.PRODUCTSENTIMENT,
				&singleItem.PRODUCTSENTIMENTMOOD,
				&singleItem.PRODUCTUSERCOMMENTS,
				&singleItem.PRODUCTRATING,
				&singleItem.PRODUCTDATEPUBLISHED,
				&singleItem.PRODUCTPRICE,
				&singleItem.PRODUCTLOCATIONCOUNT,
				&singleItem.PRODUCTINGREDIENTS,
				&singleItem.PRODUCTCATEGORY,
				&singleItem.PRODUCTIMAGE1,
				&singleItem.PRODUCTIMAGE2,
			); errX == nil {
				// Meaning we have one result
				pRecommendationLists = append(pRecommendationLists, singleItem)
			} else if errX == sql.ErrNoRows {
				// Meaning no result
				errZ := database.DB.QueryRow(`
					SELECT
					all_products.title AS productname, 
					all_products.product_name_clean_url AS productGUID, 
					product_categories.category AS category FROM all_products 
					JOIN product_categories ON all_products.category = product_categories.id 
					 WHERE all_products.product_name_clean_url = ?
				`, &j.PRODUCTGUID).Scan(&x.PRODUCTNAME, &x.PRODUCTGUID, &x.PRODUCTCATEGORY)
				if errZ != nil {
					return nil, nil, errZ
				}
				noCompetitionLists = append(noCompetitionLists, x)

			} else {
				return nil, nil, errX
			}

			/////////////////////////////////////////////////////////////////////
		}
		defer row.Close()
		return pRecommendationLists, noCompetitionLists, nil
	}
}

func deleteRating(rid, pid, user string) (int, error) {

	// We need to remember to generate sku and mpn
	stmt, err := database.DB.Prepare(`
						DELETE FROM product_review
						WHERE id = ? AND
						product_id = ?
						AND user = ?
								`)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	res, err := stmt.Exec(rid, pid, user)

	if err != nil {
		return 0, errors.New(err.Error())
	}

	lid, err := res.RowsAffected()
	return int(lid), nil

}

func getProductToEdit(pid string) ([]productListPageEdit, error) {
	var productPage = []productListPageEdit{}
	var (
		singleItem productListPageEdit
	)
	row, err := database.DB.Query(`
		SELECT all_products.id AS product_ID, 
		all_products.title AS productname,
		all_products.product_name_clean_url AS productGUID, 
		all_products.about AS description, 
		all_products.manufacturer AS manufacturer, 
		all_products.address AS address, 
		SUM(product_review.likes) AS likes, 
		SUM(product_review.dislikes) AS dislikes, 
		CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'trending_up' ELSE 'trending_down' END AS trend, 
		CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'Up' ELSE 'Down' END AS trend_direction, 
		CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'sentiment_very_satisfied' ELSE 'sentiment_very_dissatisfied' END AS sentiment, 
		CASE WHEN (product_review.likes) > (product_review.dislikes) THEN 'Good' ELSE 'Bad' END AS sentiment_mood, 
		COUNT(product_review.user_comments) AS usercomments, 
		AVG(product_review.rating) AS rating,
		product_review.author AS author,
		UNIX_TIMESTAMP(product_review.date) AS datePublished, 
		all_products.sku AS sku, 
		all_products.mpn AS mpn,
		all_products.price AS price, 
		COUNT(user_location_lat) + COUNT(user_location_lon) AS locationcount, 
		all_products.competitor_1 AS firstCompetition, 
		all_products.competitor_2 AS secondCompetition,
		all_products.ingredients AS ingredients, 
		product_categories.category AS category, 
		CONCAT('https://asknigeria.com.ng/assets/brands/images/281x224/', SUBSTR(all_products.product_image_1,8)) AS productImage_1, 
		CONCAT('https://asknigeria.com.ng/assets/brands/images/750x224/', SUBSTR(all_products.product_image_2,8)) AS productImage_2 
		FROM all_products 
		JOIN product_review ON all_products.id = product_review.product_id
		JOIN product_categories ON all_products.category = product_categories.id  
		WHERE 
		all_products.about <> '' 
		AND all_products.about IS NOT NULL
		AND all_products.product_name_clean_url = ?
		AND all_products.manufacturer IS NOT NULL AND 
		all_products.address IS NOT NULL AND
		all_products.ingredients IS NOT NULL AND
		all_products.product_image_1 IS NOT NULL AND
		all_products.product_image_2 IS NOT NULL AND
		all_products.price IS NOT NULL
		GROUP BY all_products.id ORDER BY rating DESC
	`, pid)
	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&singleItem.PRODUCTID,
				&singleItem.PRODUCTNAME,
				&singleItem.PRODUCTGUID,
				&singleItem.PRODUCTDESCRIPTION,
				&singleItem.PRODUCTMANUFACTURER,
				&singleItem.PRODUCTMANUFACTURERADDR,
				&singleItem.PRODUCTLIKES,
				&singleItem.PRODUCTDISLIKES,
				&singleItem.PRODUCTTREND,
				&singleItem.PRODUCTTRENDDIRECTION,
				&singleItem.PRODUCTSENTIMENT,
				&singleItem.PRODUCTSENTIMENTMOOD,
				&singleItem.PRODUCTUSERCOMMENTS,
				&singleItem.PRODUCTRATING,
				&singleItem.PRODUCTAUTHOR,
				&singleItem.PRODUCTDATEPUBLISHED,
				&singleItem.PRODUCTSKU,
				&singleItem.PRODUCTMPN,
				&singleItem.PRODUCTPRICE,
				&singleItem.PRODUCTLOCATIONCOUNT,
				&singleItem.PRODUCTFIRSTCOMPETITION,
				&singleItem.PRODUCTSECONDCOMPETITION,
				&singleItem.PRODUCTINGREDIENTS,
				&singleItem.PRODUCTCATEGORY,
				&singleItem.PRODUCTIMAGE1,
				&singleItem.PRODUCTIMAGE2,
			)
			if err != nil {
				return nil, err
			}
			productPage = append(productPage, singleItem)
		}
		defer row.Close()
	}
	return productPage, nil
}

func editProductDB(items MyformE) (*MyformE, error) {
	// We need to remember to generate sku and mpn
	price, _ := strconv.ParseFloat(items.PRICE, 64)
	priceF := float64(price)
	about := strings.TrimSpace(items.ABOUT)
	// about = strconv.Quote(about)
	ingredients := strings.TrimSpace(items.INGREDIENTS)
	// ingredients = strconv.Quote(ingredients)
	log.Println(items.PRODUCTID)
	updateAllPr, err := database.DB.Prepare(`update all_products
						SET 
						title = ?, 
						category = ?, 
						competitor_1 = ?, 
						competitor_2 = ?,
						about = ?, 
						manufacturer = ?, 
						address = ?, 
						ingredients = ?, 
						updated = ?, 
						price = ?
						WHERE id = ?`)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	tx, err := database.DB.Begin()

	_, err = tx.Stmt(updateAllPr).Exec(
		items.PRODUCTNAME,
		items.CATEGORIES,
		items.COMPETITORS[0],
		items.COMPETITORS[1],
		about,
		items.MANUFACTURER,
		items.MANUFACTURERADDRESS,
		ingredients,
		0,
		priceF,
		items.PRODUCTID,
	)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer updateAllPr.Close()

	if err := tx.Commit(); err != nil {
		return nil, errors.New(err.Error())
	}

	return &items, nil
}
