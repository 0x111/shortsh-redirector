package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/short-sh/shortsh-backend/models"
	"github.com/short-sh/shortsh-redirector/utils"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"runtime"
	"strings"
)

var engine *xorm.Engine

func main() {
	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	var err error
	// read config
	viper.SetConfigName("config")    // name of config file (without extension)
	viper.AddConfigPath("./_config") // optionally look for config in the working directory
	viper.AddConfigPath(".")         // optionally look for config in the working directory
	err = viper.ReadInConfig()       // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	engine, err = xorm.NewEngine("mysql", viper.GetString("mysql_dsn"))

	if err != nil {
		log.Fatalf("We could not connect to the database %v\n", err)
	}

	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)

	e := echo.New()

	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{}))

	e.GET("/:shortId", func(c echo.Context) error {
		shortID := c.Param("shortId")

		// if shortID has + at the end, it means we need the stats url for the current shortid
		if strings.HasSuffix(shortID, "+") {
			return c.Redirect(http.StatusMovedPermanently, viper.GetString("main_site")+"/"+shortID)
		}

		var shortURL = models.Url{ShortId: shortID}
		has, err := engine.Get(&shortURL)

		if has && err == nil {
			utils.WriteVisitorsData(engine, c, &shortURL)
			return c.Redirect(http.StatusMovedPermanently, shortURL.Url)
		}

		return c.Redirect(http.StatusMovedPermanently, "https://short.sh")
	})

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "https://short.sh")
	})

	e.Logger.Fatal(e.Start(":1324"))
}
