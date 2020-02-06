package main

import (
	serverConfig "adv/config"
	"adv/middleware"
	"adv/controller"
	"flag"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"os"
)
import _ "github.com/jinzhu/gorm/dialects/sqlite"

var config *serverConfig.Config

var (
	debug      bool
	configFile string
)

func main() {
	flag.StringVar(&configFile, "f", "", "config file")
	flag.BoolVar(&debug, "d", false, "enable debugging")
	if !argsChecked() {
		os.Exit(-1)
	}
	// init config
	config = serverConfig.LoadConfig(serverConfig.FromJsonFile(configFile), serverConfig.FromEnv)

	// open database
	db, dbError := gorm.Open("sqlite3", "file:"+config.DbFile)
	if dbError != nil {
		fmt.Printf("database open error: %s", dbError)
		os.Exit(-1)
	}
	defer db.Close()
	if debug {
		// gorm debug mode
		db.LogMode(true)
	}

	s := echo.New()
	s.Use(session.Middleware(sessions.NewCookieStore([]byte(config.CookieSecret))))
	s.Use(middleware.CheckLoginMiddleware)
	middleware.AddExcludeURI("/", "/index.html")

	s.Static("/", config.StaticDir)

	s.GET("/", func(c echo.Context) error {
		return c.Redirect(301, "/index.html")
		//return c.String(http.StatusOK, "Hello, World!")
	})
	controller.ApiService(s, db)
	s.Logger.Fatal(s.Start(":1323"))
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
