package books

import (
	"fmt"
	"go-playground/apps/accounts"
)

type Book struct {
	Id         int64         `json:"id"`
	Title      string        `json:"title"`
	TitleImage string        `json:"titleImage"`
	AuthorId   int64         `json:"authorId"`
	Author     accounts.User `pg:"fk:author_id",json:"id"`
}

func (b *Book) str() string {
	return fmt.Sprintf("Book<%d %s>", b.Id, b.Title)
}
