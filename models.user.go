// models.user.go

package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/brandsnigeria/webapp/database"
	"github.com/gin-gonic/gin"
)

type ReviewUser struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"-"`
}

// For this demo, we're storing the user list in memory
// We also have some users predefined.
// In a real application, this list will most likely be fetched
// from a database. Moreover, in production settings, you should
// store passwords securely by salting and hashing them instead
// of using them as we're doing in this demo
var userList = []ReviewUser{}

func getUsers() {

	var (
		user ReviewUser
	)
	row, err := database.DB.Query("select firstname, lastname, password, username from users")
	if err != nil {
		log.Println(err.Error())
	} else {
		for row.Next() {
			err = row.Scan(&user.Firstname, &user.Lastname, &user.Username, &user.Password)
			userList = append(userList, user)
			if err != nil {
				log.Println(err.Error())
			}
		}
		defer row.Close()
	}
}

// Register a new user with the given username and password
func registerNewUser(firstname, lastname, username, password string) (*ReviewUser, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("The username isn't available")
	}
	u := ReviewUser{Firstname: firstname, Lastname: lastname, Username: username, Password: password}
	stmt, err := database.DB.Prepare(`insert into users (firstname, lastname, username, password, super_admin) 
									values(?,?,?,?,?);`)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	_, err = stmt.Exec(firstname, lastname, username, password, 0)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer stmt.Close()

	return &u, nil
}

// Check if the supplied username is available
func isUsernameAvailable(username string) bool {
	var (
		numCount int
	)
	row, err := database.DB.Query("select COUNT(*) as count from users where username = ?", username)
	numCount = checkCount(row)
	checkErr(err)
	if numCount > 0 {
		return false
	}
	return true
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		render(c, gin.H{"title": "Server Error", "message": http.StatusServiceUnavailable}, "500.tmpl")
	}
}

func isUserValid(username, password string) bool {
	var (
		numCount int
	)
	row, err := database.DB.Query(`select count(*) AS count from users 
				WHERE username = ? AND password = ?`, username, password)
	numCount = checkCount(row)
	checkErr(err)
	if numCount > 0 {
		return true
	}
	return false
}
