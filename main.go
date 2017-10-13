package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	tablename = "bookstate"
	conn, _   = dbr.Open("mysql", "username:api@/book_api", nil)
	sess      = conn.NewSession(nil)
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	// Routing
	e.GET("/book/pullbookinfo/:title", PullBookInfo)
	e.GET("/book/pullbooksinfo/all", PullBooksInfo)
	e.GET("/book/pullbooksinfo/keyword/:key", PullBooksKeyword)
	e.PUT("/book/updatebookinfo", UpdateBookInfo)
	e.POST("/book/postbookinfo", PostBookInfo)
	e.DELETE("/book/deletebookinfo/:id", DeleteBookInfo)

	e.Logger.Fatal(e.Start(":9090"))
}
