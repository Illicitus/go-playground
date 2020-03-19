package accounts

import (
	"encoding/json"
	"go-playground/core"
	"net/http"
)

func (u *User) decodeRequestData(w http.ResponseWriter, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(u)
	if core.JsonErrorHandler400(w, err) {
		return err
	}
	return nil
}
