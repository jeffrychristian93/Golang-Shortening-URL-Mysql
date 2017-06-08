package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


type Redirect struct {
	Id 		int
	Slug 	string 	`db:"slug" form:"slug"`
	Url  	string	`db:"url" form:"url"`
}


var db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/db-redirects")


func getBySlug(c *gin.Context){
	var redirect Redirect

	slug := c.Param("slug")
	row := db.QueryRow("select id, slug, url from redirect where slug = ?;", slug)
	
	err = row.Scan(&redirect.Id, &redirect.Slug, &redirect.Url)
	if err != nil {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("301 Found"),
			"url": fmt.Sprintf("Location: %s",redirect.Url),
		})
	}
}

func generateSlug() string {
	var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyz")
	s := make([]rune, 6)
	for i := range s {
		s[i] = chars[rand.Intn(len(chars))]
	}
	return string(s)
}


func add(c *gin.Context) {
	Slug := generateSlug()
	Url := c.PostForm("url")

	stmt, err := db.Prepare("insert into redirect (slug, url) values(?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(Slug, Url)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("201 Created"),
		"url": fmt.Sprintf("Location: http://domain.com/%s",Slug),
	})
	
}


func createTable(){
	stmt, err := db.Prepare("CREATE TABLE redirect (id int NOT NULL AUTO_INCREMENT, Slug varchar(40), Url varchar(40), PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Redirect Table successfully migrated....")
	}
}


func main() {
	createTable();
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	
	router := gin.Default()
	router.GET("/:slug", getBySlug)
	router.POST("/create", add)
	router.Run(":8000")
}