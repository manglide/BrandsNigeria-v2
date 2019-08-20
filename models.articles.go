// models.article.go

package main

import (
	"errors"

	"time"

	"github.com/brandsnigeria/webapp/database"
)

func getSingleArticle(guid string) ([]blog, error) {
	var lb = []blog{}
	var (
		singleItem blog
	)
	row, err := database.DB.Query(`
		SELECT id, title, content, date_published FROM blog WHERE guid = ?
	`, guid)

	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err = row.Scan(
				&singleItem.ID,
				&singleItem.Title,
				&singleItem.Content,
				&singleItem.Date,
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

func createNewArticle(title, content string) (*blog, error) {
	var (
		singleItem blog
	)
	var datetime = time.Now()
	datetime.Format(time.RFC3339)
	guid := sGUID(title)
	stmtX, errX := database.DB.Prepare(`insert into blog
				(
					title, content, date_published, author, guid
				)
				values(?,?,?,?,?);`)
	if errX != nil {
		return nil, errors.New(errX.Error())
	}
	resX, errX := stmtX.Exec(
		title, content, datetime, UserLoggedIn, guid,
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

	if errX != nil {
		return nil, errors.New(errX.Error())
	}
	return &singleItem, nil
}
