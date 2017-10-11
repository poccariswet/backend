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
	e.GET("/book/pullbookinfo/:title", Pull_Book_Info)
	e.GET("/book/pullbooksinfo/all", Pull_Books_Info)
	e.PUT("/book/updatebookinfo", Update_Book_Info)
	e.POST("/book/postbookinfo", Post_Book_Info)
	e.DELETE("/book/deletebookinfo/:id", Delete_Book_Info)

	e.Logger.Fatal(e.Start(":9090"))
}
