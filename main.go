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
	e.GET("/book/api/get/:title", PullBookInfo)
	e.GET("/book/api/get/all", PullBooksInfo)
	e.GET("/book/api/get/key/:key", PullBooksKeyword)
	e.PUT("/book/api/update", UpdateBookInfo)
	e.POST("/book/api/post", PostBookInfo)
	e.DELETE("/book/api/delete/:id", DeleteBookInfo)

	e.Logger.Fatal(e.Start(":9090"))
}
