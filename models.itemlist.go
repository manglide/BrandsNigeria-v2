// models.article.go

package main

import (
	"github.com/brandsnigeria/webapp/database"
)

var itemLists = []itemList{}
var productLists = []productList{}

type itemList struct {
	PRODUCTID   int    `json:"id"`
	PRODUCTGUID string `json:"productGUID"`
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

func getAllItemsFrontPage() ([]productList, error) {
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
