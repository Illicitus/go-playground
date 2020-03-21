package books

import (
	"errors"
	"go-playground/core"
	"gopkg.in/validator.v2"
	"net/http"
)

func (b *CreteUpdateBookSchema) validate(w http.ResponseWriter) error {

	if err := validator.Validate(b); err != nil {
		core.JsonErrorHandler400(w, err)
		return err
	}

	db := core.GetDb()
	status, err := db.Model(&BookTitleImage{}).Where("id = ?", b.TitleImageId).Exists()
	if core.JsonErrorHandler500(w, err) {
		return err
	}

	if !status {
		core.JsonErrorHandler400(w, errors.New("title image id doesn't exist"))
		return errors.New("")
	}
	return nil
}

func (b *Book) decodeAndValidate(w http.ResponseWriter, r *http.Request) error {
	var createUpdateBook CreteUpdateBookSchema

	if err := core.DecodeRequestData(&createUpdateBook, w, r); err != nil {
		return err
	}

	if err := createUpdateBook.validate(w); err != nil {
		return err
	}

	// Update book data
	b.Title = createUpdateBook.Title
	b.TitleImageId = createUpdateBook.TitleImageId

	return nil
}

func (bc *CreteUpdateBookCommentSchema) validate(w http.ResponseWriter) error {
	if err := validator.Validate(bc); err != nil {
		core.JsonErrorHandler400(w, err)
		return err
	}
	return nil
}

func (bc *BookComment) decodeAndValidate(w http.ResponseWriter, r *http.Request) error {
	var createBookComment CreteUpdateBookCommentSchema

	if err := core.DecodeRequestData(&createBookComment, w, r); err != nil {
		return err
	}

	if err := createBookComment.validate(w); err != nil {
		return err
	}

	// Update book comment data
	bc.Message = createBookComment.Message

	return nil
}
