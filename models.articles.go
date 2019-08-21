// models.article.go

package main

import (
	"errors"

	"log"
	"time"

	"github.com/brandsnigeria/webapp/database"
)

func getSingleArticle(id int, guid string) ([]blog, error) {
	var lb = []blog{}
	var (
		singleItem blog
	)
	row, err := database.DB.Query(`
		SELECT id, title, content, guid, date_published,
		CASE WHEN imageloc = ''
            THEN ''
            ELSE 
            CONCAT('https://images.brandsnigeria.com.ng/', imageloc)
       	END AS imageloc 
		FROM blog WHERE id = ? AND guid = ?
	`, id, guid)

	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&singleItem.ID,
				&singleItem.Title,
				&singleItem.Content,
				&singleItem.GUID,
				&singleItem.Date,
				&singleItem.IMAGE,
			)
			if err != nil {
				return nil, err
			}
			lb = append(lb, singleItem)
		}
		defer row.Close()
	}
	return lb, nil

}

func createNewArticle(title, content, imageloc string) (*blog, error) {
	var (
		singleItem blog
	)
	var datetime = time.Now()
	datetime.Format(time.RFC3339)
	guid := sGUID(title)
	log.Println(imageloc)
	stmtX, errX := database.DB.Prepare(`insert into blog
				(
					title, content, date_published, author, guid, imageloc
				)
				values(?,?,?,?,?,?);`)
	if errX != nil {
		return nil, errors.New(errX.Error())
	}
	resX, errX := stmtX.Exec(
		title, content, datetime, UserLoggedIn, guid, imageloc,
	)

	if errX != nil {
		return nil, errors.New(errX.Error())
	}

	defer stmtX.Close()

	lid, errX := resX.LastInsertId()

	singleItem.ID = int(lid)
	singleItem.Date = datetime.String()
	singleItem.GUID = guid
	singleItem.Title = title
	singleItem.IMAGE = imageloc

	if errX != nil {
		return nil, errors.New(errX.Error())
	}
	return &singleItem, nil
}
