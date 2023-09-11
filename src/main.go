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
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/article/handler"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/article/infra"
	handlerc "github.com/wataruhigasi/Katurao-Hackathon-Back/components/comment/handler"
	infrac "github.com/wataruhigasi/Katurao-Hackathon-Back/components/comment/infra"

	handlert "github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/handler"
	infrat "github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/usecase"
)

func main() {
	conn, err := initDB("mysql-server", 3306, "katurao", "admin", "password")
	if err != nil {
		log.Fatal(err)
	}

	ar := infra.NewArticleRepository(conn)
	ah := handler.NewArticleHandler(ar)

	cr := infrac.NewCommentRepository(conn)
	ch := handlerc.NewCommentHandler(cr)

	tr := infrat.NewThreadRepository(conn)
	tu := usecase.NewThreadUsecase(tr, cr)
	th := handlert.NewThreadHandler(tu)

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
