package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Book struct {
	Id     int    `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Isbn13 string `db:"isbn13"`
	State  bool   `db:"state"`
	Pic    string `db:"pic"`
}

func Pull_Book_Info(c echo.Context) error {
	var book Book
	//	id, _ := strconv.Atoi(c.Param("id"))
	title := c.Param("title")
	sess.Select("*").From(tablename).Where("title = ?", title).Load(&book)

	return c.JSON(http.StatusOK, book)
}

func Pull_Books_Info(c echo.Context) error {
	var books []Book
	sess.Select("*").From(tablename).Load(&books)

	return c.JSON(http.StatusOK, books)
}

func Update_Book_Info(c echo.Context) error {
	book := new(Book)
	if err := c.Bind(book); err != nil {
		return err
	}

	attrsMap := map[string]interface{}{"id": book.Id, "title": book.Title, "author": book.Author, "isbn13": book.Isbn13, "state": book.State, "pic": book.Pic}
	sess.Update(tablename).SetMap(attrsMap).Where("id = ?", book.Id).Exec()

	return c.NoContent(http.StatusOK)
}

func Post_Book_Info(c echo.Context) error {
	book := new(Book)
	if err := c.Bind(book); err != nil {
		return err
	}

	sess.InsertInto(tablename).Columns("title", "author", "isbn13", "state", "pic").Values(book.Title, book.Author, book.Isbn13, book.State, book.Pic).Exec()

	return c.NoContent(http.StatusOK)
}

func Delete_Book_Info(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	sess.DeleteFrom(tablename).Where("id = ?", id).Exec()

	return c.NoContent(http.StatusNoContent)
}
