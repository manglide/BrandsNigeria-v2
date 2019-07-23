// models.article.go

package main

import (
	"math/rand"
	"time"

	"errors"

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

type competitors struct {
	PRODUCTID   int    `json:"id"`
	PRODUCTGUID string `json:"productguid"`
	PRODUCTNAME string `json:"productname"`
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

func getCompetitors() ([]competitors, error) {
	var competitorsLists = []competitors{}
	var (
		c competitors
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
			competitorsLists = append(competitorsLists, c)
		}
		defer row.Close()
	}
	return competitorsLists, nil
}

func insertComments(pid, cat, username, comments, rating, sentiment, latitude, longitude string) (int, error) {
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
		pid, cat, like, dislike, rating, comments, latitude, longitude, datetime, username, username,
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

}
