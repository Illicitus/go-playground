package books

import (
	"errors"
	"go-playground/core"
	"gopkg.in/validator.v2"
	"net/http"
)

func (b *Book) validate(w http.ResponseWriter) error {
	if err := validator.Validate(b); err != nil {
		core.JsonErrorHandler400(w, err)
		return err
	}

	db := core.GetDb()
	status, err := db.Model(&BookTitleImage{}).Where("id = ?", b.TitleImage.Id).Exists()
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

	if err := core.DecodeRequestData(b, w, r); err != nil {
		return err
	}

	if err := b.validate(w); err != nil {
		return err
	}
	return nil
}
