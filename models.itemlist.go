// models.article.go

package main

import (
	"github.com/brandsnigeria/webapp/database"
)

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
	PRODUCTFULLCOMMENTS      string `json:"productCOMMENTS"`
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
	LIKES         string `json:"likes"`
	DISLIKE       string `json:"dislike"`
	PRODUCTRATING string `json:"rate"`
	USERCOMMENTS  string `json:"usercomments"`
	LATITUDE      string `json:"latitude"`
	LONGITUDE     string `json:"longitude"`
	DATEPUBLISHED string `json:"datePublished"`
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

func getAllComments(productID string) ([]commentList, error) {
	var commentLists = []commentList{}
	var (
		comment commentList
	)

	row, err := database.DB.Query(`
			SELECT 
				likes AS likes, 
				dislikes AS dislike, 
				rating AS rate, 
	            user_comments AS comment, 
	            user_location_lat AS latitude, 
	            user_location_lon AS longitude
	            date AS datePublished 
            FROM product_review WHERE product_id = ?
	`, productID)
	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&comment.LIKES,
				&comment.DISLIKE,
				&comment.PRODUCTRATING,
				&comment.USERCOMMENTS,
				&comment.LATITUDE,
				&comment.LONGITUDE,
				&comment.DATEPUBLISHED,
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
				&singleItem.PRODUCTFULLCOMMENTS,
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
