package core

import (
	"encoding/json"
	"net/http"
)

func DecodeRequestData(o interface{}, w http.ResponseWriter, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(o)
	if JsonErrorHandler400(w, err) {
		return err
	}
	return nil
}

//func decodeUrlEncodedValue(w http.ResponseWriter, s string) (string, error) {
//	decodedValue, err := url.QueryUnescape(s)
//	if JsonErrorHandler500(w, err) {
//		return "", err
//	}
//	return decodedValue, nil
//}
