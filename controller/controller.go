package controller

import (
	"adv/form"
	"adv/middleware"
	"adv/service"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"
	"strings"
)

func responseData(value interface{}) map[string]interface{} {
	switch value.(type) {
	case bool:
		return map[string]interface{}{"success": value}
	default:
		return map[string]interface{}{"data": value,}
	}
}

func ApiService(s *echo.Echo, db *gorm.DB) {
	postService := service.NewPostService(db)

	// get posts list
	s.GET("/api/posts", func(c echo.Context) error {
		pager := new(form.Pager)
		c.Bind(pager)
		return c.JSON(http.StatusOK, responseData(postService.GetPostList(pager)))
	})

	// set adv json for specified post
	s.POST("/api/posts/ads", func(c echo.Context) error {
		postAdvInfo := new(form.PostAdvInfo)
		c.Bind(postAdvInfo)
		return c.JSON(http.StatusOK, responseData(postService.SetPostAdvJSON(postAdvInfo)))
	})

	// get adv json for specified post
	middleware.AddExcludeURI("/api/posts/adv")
	s.GET("/api/posts/adv", func(c echo.Context) error {
		refURI := strings.Trim(c.Request().Referer(), "/")
		uriParts := strings.Split(refURI, "/")
		mdFileName := uriParts[len(uriParts)-1] + ".md"
		return c.JSON(http.StatusOK, responseData(postService.GetPostAds(mdFileName)))
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		hostRegex, err := regexp.Compile("^https?://[^/]+/?")
		if err != nil {
			panic(err)
		}
		return func(c echo.Context) error {
			host := hostRegex.Find([]byte(c.Request().Referer()))
			if matched, _ := regexp.Match("\\bzvz.im\\b", host); matched {
				c.Response().Header().Set("Access-Control-Allow-Origin", strings.Trim(string(host), "/"))
			}
			return next(c)
		}
	})

	userService := service.NewUserService(db)

	middleware.AddExcludeURI("/api/users/login")
	s.POST("/api/users/login", func(c echo.Context) error {
		loginForm := new(form.LoginForm)
		c.Bind(loginForm)
		user, err := userService.Login(loginForm.Username, loginForm.Password)
		if user != nil {
			// save to session
			sess, _ := session.Get("session", c)
			sess.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   86400,
				HttpOnly: true,
			}
			sess.Values["userId"] = user.Id
			err := sess.Save(c.Request(), c.Response())
			if err != nil {
				return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
			}
			return c.JSON(http.StatusOK, responseData(map[string]string{"username": user.Username, "name": user.Name}))
		} else {
			var errMessage string
			if err == nil {
				errMessage = "user not found"
			} else {
				errMessage = err.Error()
			}
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: errMessage}
		}
	})
}
