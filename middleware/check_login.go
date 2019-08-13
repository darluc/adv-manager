package middleware

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckLoginMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uri := c.Request().RequestURI
		if uri != "/api/users/login" && uri != "/" && uri != "/index.html" {
			sess, _ := session.Get("session", c)
			if userId := sess.Values["userId"]; userId != nil {
				if userId.(int) > 0 {
					return next(c)
				}
			}
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "user not login"}
		}
		return next(c)
	}
}
