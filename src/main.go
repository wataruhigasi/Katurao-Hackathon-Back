package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	article_handler "github.com/wataruhigasi/Katurao-Hackathon-Back/components/article/handler"
	article_infra "github.com/wataruhigasi/Katurao-Hackathon-Back/components/article/infra"

	thread_handler "github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/handler"
	thread_infra "github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/infra"

	keijiban_rakugaki_handler "github.com/wataruhigasi/Katurao-Hackathon-Back/components/keijiban_rakugaki/handler"
	keijiban_rakugaki_infra "github.com/wataruhigasi/Katurao-Hackathon-Back/components/keijiban_rakugaki/infra"

	thread_rakugaki_handler "github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread_rakugaki/handler"
	thread_rakugaki_infra "github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread_rakugaki/infra"
)

func main() {
	conn, err := initDB("mysql-server", 3306, "katurao", "admin", "password")
	if err != nil {
		log.Fatal(err)
	}

	ar := article_infra.NewRepo(conn)
	ah := article_handler.New(ar)

	tr := thread_infra.NewRepo(conn)
	th := thread_handler.New(tr)

	krr := keijiban_rakugaki_infra.NewRepo(conn)
	krh := keijiban_rakugaki_handler.New(krr)

	trr := thread_rakugaki_infra.NewRepo(conn)
	trh := thread_rakugaki_handler.New(trr)

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", hello)
	e.GET("/articles", ah.GetAll)
	e.POST("/article", ah.Create)
	e.PATCH("/article/:article_id/position", ah.ChangePosition)

	e.GET("/threads", th.GetAll)
	e.POST("/thread", th.Create)
	e.PATCH("/thread/:thread_id/position", th.ChangePosition)

	e.GET("/keijiban/rakugakis", krh.GetAll)
	e.POST("/keijiban/rakugaki", krh.Create)

	e.GET("/thread/:thread_id/rakugakis", trh.GetAll)
	e.POST("/thread/:thread_id/rakugaki", trh.Create)

	log.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func initDB(host string, port uint16, dbName, username, password string) (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.Addr = net.JoinHostPort(host, strconv.Itoa(int(port)))
	cfg.DBName = dbName
	cfg.User = username
	cfg.Passwd = password
	cfg.ParseTime = true

	connector, err := mysql.NewConnector(cfg)
	if err != nil {
		return nil, fmt.Errorf("new connector: %w", err)
	}

	return sql.OpenDB(connector), nil
}
