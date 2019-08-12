package main

import (
	"adv/formdata"
	"adv/service"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"regexp"
	"strings"
)
import _ "github.com/jinzhu/gorm/dialects/sqlite"

var db *gorm.DB
var server *echo.Echo

func main() {
	var dbError error
	//@todo read db config
	db, dbError = gorm.Open("sqlite3", "file:/Users/darluc/Downloads/blog_manage")
	if dbError != nil {
		fmt.Printf("database open error: %s", dbError)
		os.Exit(-1)
	}
	defer db.Close()

	server = echo.New()
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	startupAPIServer()
	server.Logger.Fatal(server.Start(":1323"))
}

func startupAPIServer() {
	ps := service.NewPostService(db)

	// get posts list
	server.GET("/api/posts", func(c echo.Context) error {
		pager := new(formdata.Pager)
		c.Bind(pager)
		return c.JSON(http.StatusOK, ps.GetPostList(pager))
	})

	// set adv json for specified post
	server.POST("/api/posts/ads", func(c echo.Context) error {
		postAdvInfo := new(formdata.PostAdvInfo)
		c.Bind(postAdvInfo)
		return c.JSON(http.StatusOK, ps.SetPostAdvJSON(postAdvInfo))
	})

	// get adv json for specified post
	server.GET("/api/posts/adv", func(c echo.Context) error {
		host := c.Request().Host
		if matched, _ := regexp.Match("\\bzvz.im\\b", []byte(host)); matched {
			c.Response().Header().Set("Access-Control-Allow-Origin", "https://"+host)
		}
		refURI := strings.Trim(c.Request().Referer(), "/")
		uriParts := strings.Split(refURI, "/")
		mdFileName := uriParts[len(uriParts)-1] + ".md"
		return c.JSON(http.StatusOK, ps.GetPostAds(mdFileName))
	})
}

func serveFrontendFiles() {

}
