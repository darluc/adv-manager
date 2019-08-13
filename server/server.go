package main

import (
	serverConfig "adv/config"
	"adv/middleware"
	"flag"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"os"
)
import _ "github.com/jinzhu/gorm/dialects/sqlite"

var db *gorm.DB
var server *echo.Echo
var config *serverConfig.Config

var configFile string

func main() {
	flag.StringVar(&configFile, "f", "", "config file")
	if !argsChecked() {
		os.Exit(-1)
	}
	// init config
	config = serverConfig.LoadConfig(serverConfig.FromJsonFile(configFile))

	// open database
	var dbError error
	db, dbError = gorm.Open("sqlite3", "file:"+config.DbFile)
	if dbError != nil {
		fmt.Printf("database open error: %s", dbError)
		os.Exit(-1)
	}
	defer db.Close()

	server = echo.New()
	server.Use(session.Middleware(sessions.NewCookieStore([]byte(config.CookieSecret))))
	server.Use(middleware.CheckLoginMiddleware)

	server.Static("/", config.StaticDir)

	server.GET("/", func(c echo.Context) error {
		return c.Redirect(301, "/index.html")
		//return c.String(http.StatusOK, "Hello, World!")
	})
	startupAPIService()
	server.Logger.Fatal(server.Start(":1323"))
}

func argsChecked() bool {
	flag.Parse()
	if configFile == "" {
		fmt.Fprint(os.Stderr, "config file not specified")
		return false
	}
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "file not exists: %s", configFile)
		return false
	}
	return true
}
