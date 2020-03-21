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
