package main

import (
	"adv/formdata"
	"adv/middleware"
	"adv/service"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
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
	// @todo get cookie secret string from config
	server.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	server.Use(middleware.CheckLoginMiddleware)

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	startupAPIServer()
	server.Logger.Fatal(server.Start(":1323"))
}

//@todo: route middleware

//@todo: global error handler ?
func startupAPIServer() {
	postService := service.NewPostService(db)

	// get posts list
	server.GET("/api/posts", func(c echo.Context) error {
		pager := new(formdata.Pager)
		c.Bind(pager)
		return c.JSON(http.StatusOK, postService.GetPostList(pager))
	})

	// set adv json for specified post
	server.POST("/api/posts/ads", func(c echo.Context) error {
		postAdvInfo := new(formdata.PostAdvInfo)
		c.Bind(postAdvInfo)
		return c.JSON(http.StatusOK, postService.SetPostAdvJSON(postAdvInfo))
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
		return c.JSON(http.StatusOK, postService.GetPostAds(mdFileName))
	})

	userService := service.NewUserService(db)

	server.POST("/api/users/login", func(c echo.Context) error {
		loginForm := new(formdata.LoginForm)
		c.Bind(loginForm)
		user, err := userService.Login(loginForm.Username, loginForm.Password)
		if user != nil {
			// save to session
			sess, _ := session.Get("session", c)
			sess.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   86400,
				HttpOnly: true,
				Secure:   true,
			}
			sess.Values["userId"] = user.Id
			err := sess.Save(c.Request(), c.Response())
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, map[string]string{"username": user.Username, "name": user.Name})
		} else {
			return err
		}
	})
}
