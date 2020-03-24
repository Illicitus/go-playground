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
	Id           int64            `json:"id"`
	Title        string           `validate:"nonzero",json:"title"`
	TitleImageId int64            `pg:"fk:titleImage_id",json:"id"`
	TitleImage   BookTitleImage   `json:"titleImage"`
	AuthorId     int64            `json:"authorId"`
	Author       accounts.User    `pg:"fk:author_id",json:"id"`
	Likes        []*accounts.User `pg:"many2many:book_likes"`
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

func listBooksByAuthorId(w http.ResponseWriter, authorId int64) ([]Book, error) {
	var books []Book
	db := core.GetDb()

	err := db.
		Model(&books).
		Where("author_id = ?", authorId).
		ColumnExpr("book.*").
		ColumnExpr("u.id AS author__id, u.name AS author__name, u.email AS author__email").
		ColumnExpr("t.id AS title_image__id, t.source AS title_image__source, t.thumbnail AS title_image__thumbnail").
		Join("JOIN users AS u ON u.id = book.author_id").
		Join("JOIN book_title_images AS t ON t.id = book.title_image_id").
		Order("id ASC").
		Select()

	if core.JsonErrorHandler500(w, err) {
		return books, err
	}
	return books, nil
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

type BookLikes struct {
	BookId int64
	UserId int64
}

func LikeOrDislike(w http.ResponseWriter, bui int64, uid int64) error {
	db := core.GetDb()

	if exists, err := db.Model(&BookLikes{}).Where("book_id = ?", bui).Where("user_id = ?", uid).Exists(); core.JsonErrorHandler500(w, err) {
		return err
	} else if exists {
		if _, err := db.Model(&BookLikes{}).Where("book_id = ?", bui).Where("user_id = ?", uid).Delete(); core.JsonErrorHandler500(w, err) {
			return err
		}
	} else {
		if err := db.Insert(&BookLikes{BookId: bui, UserId: uid}); core.JsonErrorHandler500(w, err) {
			return err
		}
	}
	return nil
}

type BookComment struct {
	Id       int64         `json:"id"`
	Message  string        `validate:"nonzero",json:"message"`
	BookId   int64         `pg:"fk:book_id",json:"id"`
	Book     Book          `json:"book"`
	AuthorId int64         `json:"authorId"`
	Author   accounts.User `pg:"fk:author_id",json:"id"`
}

func (bc *BookComment) str() string {
	return fmt.Sprintf("BookComment<%d %s>", bc.Id, bc.Book.Title)
}

func (bc *BookComment) createBookComment(w http.ResponseWriter) error {
	db := core.GetDb()
	err := db.Insert(bc)

	if core.JsonErrorHandler500(w, err) {
		return err
	}
	return nil
}

func (bc *BookComment) getBookCommentById(w http.ResponseWriter, id int64) error {
	db := core.GetDb()

	err := db.
		Model(bc).
		Where("book_comment.id = ?", id).
		ColumnExpr("book_comment.*").
		ColumnExpr("u.id AS author__id, u.name AS author__name, u.email AS author__email").
		Join("JOIN users AS u ON u.id = book_comment.author_id").
		Order("id ASC").
		Select()

	if core.JsonErrorHandler500(w, err) {
		return err
	}
	return nil
}

func listBookCommentsByBookId(w http.ResponseWriter, bookId int64) ([]BookComment, error) {
	var comments []BookComment
	db := core.GetDb()

	err := db.
		Model(&comments).
		Where("book_id = ?", bookId).
		ColumnExpr("book_comment.*").
		ColumnExpr("u.id AS author__id, u.name AS author__name, u.email AS author__email").
		Join("JOIN users AS u ON u.id = book_comment.author_id").
		Order("id ASC").
		Select()

	if core.JsonErrorHandler500(w, err) {
		return comments, err
	}
	return comments, nil
}
