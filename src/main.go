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
	comment_handler "github.com/wataruhigasi/Katurao-Hackathon-Back/components/comment/handler"
	comment_infra "github.com/wataruhigasi/Katurao-Hackathon-Back/components/comment/infra"

	thread_handler "github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/handler"
	thread_infra "github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/infra"
	thread_usecase "github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/usecase"
)

func main() {
	conn, err := initDB("mysql-server", 3306, "katurao", "admin", "password")
	if err != nil {
		log.Fatal(err)
	}

	ar := article_infra.NewRepo(conn)
	ah := article_handler.New(ar)

	cr := comment_infra.NewRepo(conn)
	ch := comment_handler.New(cr)

	tr := thread_infra.NewRepo(conn)
	tu := thread_usecase.New(tr, cr)
	th := thread_handler.New(tu)

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", hello)
	e.GET("/articles", ah.GetAll)
	e.POST("/article", ah.Create)

	e.GET("/thread/:thread_id/comments", ch.GetAll)
	e.POST("/thread/:thread_id/comment", ch.Create)

	e.GET("/threads", th.GetAll)
	e.POST("/thread", th.Create)

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
