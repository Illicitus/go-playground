package books

import (
	"fmt"
	"go-playground/apps/accounts"
	"go-playground/core"
	"net/http"
)

type BookTitleImage struct {
	Id        int64  `json:"id"`
	Source    string `json:"source"`
	Thumbnail string `json:"thumbnail"`
}

func (b *BookTitleImage) str() string {
	return fmt.Sprintf("BookTitleImage<%d %s %s>", b.Id, b.Source, b.Thumbnail)
}

func (b *BookTitleImage) createNewBookTitleImage(w http.ResponseWriter) error {
	db := core.GetDb()
	err := db.Insert(b)
	if core.JsonErrorHandler500(w, err) {
		return err
	}
	return nil
}

type Book struct {
	Id           int64          `json:"id"`
	Title        string         `validate:"nonzero",json:"title"`
	TitleImageId int64          `pg:"fk:titleImage_id",json:"id"`
	TitleImage   BookTitleImage `json:"titleImage"`
	AuthorId     int64          `json:"authorId"`
	Author       accounts.User  `pg:"fk:author_id",json:"id"`
}

func (b *Book) str() string {
	return fmt.Sprintf("Book<%d %s>", b.Id, b.Title)
}

func (b *Book) createNewBook(w http.ResponseWriter) error {
	db := core.GetDb()
	err := db.Insert(b)
	if core.JsonErrorHandler500(w, err) {
		return err
	}
	return nil
}

func (b *Book) getBookById(w http.ResponseWriter, id int64) error {
	db := core.GetDb()

	err := db.
		Model(b).
		Where("book.id = ?", id).
		ColumnExpr("book.*").
		ColumnExpr("u.id AS author__id, u.name AS author__name, u.email AS author__email").
		ColumnExpr("t.id AS title_image__id, t.source AS title_image__source, t.thumbnail AS title_image__thumbnail").
		Join("JOIN users AS u ON u.id = book.author_id").
		Join("JOIN book_title_images AS t ON t.id = book.title_image_id").
		Select()

	if core.JsonErrorHandler500(w, err) {
		return err
	}
	return nil
}

func (b *Book) updateBook(w http.ResponseWriter, id int64) error {
	db := core.GetDb()

	_, err := db.
		Model(b).
		Set("title = ?title").
		Set("title_image_id = ?title_image_id").
		Set("author_id = ?author_id").
		Where("id = ?", id).
		Update()

	if core.JsonErrorHandler500(w, err) {
		return err
	}
	return nil
}
