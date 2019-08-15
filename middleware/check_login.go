package middleware

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

var excludeURI []string

func AddExcludeURI(uri ...string) {
	if excludeURI == nil {
		excludeURI = make([]string, 0)
	}
	excludeURI = append(excludeURI, uri...)
}

func CheckLoginMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uri := c.Request().URL.Path
		for _, exclude := range excludeURI {
			if uri == exclude {
				return next(c)
			}
		}

		sess, _ := session.Get("session", c)
		if userId := sess.Values["userId"]; userId != nil {
			if userId.(int) > 0 {
				return next(c)
			}
		}
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "user not login"}
	}
}
