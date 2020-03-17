package books

import (
	"encoding/json"
	"go-playground/apps/accounts"
)

type BookSchema struct {
	Id         int64                       `json:"id"`
	Title      string                      `json:"title"`
	TitleImage string                      `json:"titleImage"`
	Author     *accounts.UserProfileSchema `json:"author"`
}

func serializeBookSchema(b Book) ([]byte, error) {
	return json.Marshal(BookSchema{
		Id:         b.Id,
		Title:      b.Title,
		TitleImage: b.TitleImage,
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
			Id:         v.Id,
			Title:      v.Title,
			TitleImage: v.TitleImage,
			Author: &accounts.UserProfileSchema{
				Id:    v.Author.Id,
				Name:  v.Author.Name,
				Email: v.Author.Email,
			}})
	}
	return json.Marshal(r)
}
