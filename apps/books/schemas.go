package books

import (
	"encoding/json"
	"go-playground/apps/accounts"
)

func serializeBookTitleImageSchema(b BookTitleImage) ([]byte, error) {
	return json.Marshal(BookTitleImage{
		Id:        b.Id,
		Source:    b.Source,
		Thumbnail: b.Thumbnail,
	})
}

type BookSchema struct {
	Id         int64                       `json:"id"`
	Title      string                      `json:"title"`
	TitleImage *BookTitleImage             `json:"titleImage"`
	Author     *accounts.UserProfileSchema `json:"author"`
}

func serializeBookSchema(b Book) ([]byte, error) {
	return json.Marshal(BookSchema{
		Id:    b.Id,
		Title: b.Title,
		TitleImage: &BookTitleImage{
			Id:        b.TitleImage.Id,
			Source:    b.TitleImage.Source,
			Thumbnail: b.TitleImage.Thumbnail,
		},
		Author: &accounts.UserProfileSchema{
			Id:    b.Author.Id,
			Name:  b.Author.Name,
			Email: b.Author.Email,
		},
	})
}

func serializeManyBookSchema(b []Book) ([]byte, error) {
	var r []BookSchema

	for _, v := range b {
		r = append(r, BookSchema{
			Id:    v.Id,
			Title: v.Title,
			TitleImage: &BookTitleImage{
				Id:        v.TitleImage.Id,
				Source:    v.TitleImage.Source,
				Thumbnail: v.TitleImage.Thumbnail,
			},
			Author: &accounts.UserProfileSchema{
				Id:    v.Author.Id,
				Name:  v.Author.Name,
				Email: v.Author.Email,
			}})
	}
	return json.Marshal(r)
}

type CreteUpdateBookSchema struct {
	Id           int64  `json:"id"`
	Title        string `validate:"nonzero",json:"title"`
	TitleImageId int64  `json:"titleImage"`
}

type BookCommentSchema struct {
	Id      int64                       `json:"id"`
	Message string                      `json:"message"`
	Author  *accounts.UserProfileSchema `json:"author"`
}

func serializeBookCommentSchema(bc BookComment) ([]byte, error) {
	return json.Marshal(BookCommentSchema{
		Id:      bc.Id,
		Message: bc.Message,
		Author: &accounts.UserProfileSchema{
			Id:    bc.Author.Id,
			Name:  bc.Author.Name,
			Email: bc.Author.Email,
		},
	})
}

func serializeManyBookCommentSchema(bc []BookComment) ([]byte, error) {
	var r []BookCommentSchema

	for _, v := range bc {
		r = append(r, BookCommentSchema{
			Id:      v.Id,
			Message: v.Message,
			Author: &accounts.UserProfileSchema{
				Id:    v.Author.Id,
				Name:  v.Author.Name,
				Email: v.Author.Email,
			}})
	}
	return json.Marshal(r)
}

type CreteUpdateBookCommentSchema struct {
	Id       int64  `json:"id"`
	Message  string `json:"message"`
	AuthorId int64  `json:"author"`
}
