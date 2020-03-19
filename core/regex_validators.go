package core

import "regexp"

func IsEmpty(s string) (bool, error) {
	return regexp.MatchString("^$", s)
}

func IsEmailValid(s string) (bool, error) {
	return regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", s)
}

func IsPasswordValid(s string) (bool, error) {
	return regexp.MatchString("^(?=.*[0-9])|(?=.{8,})$", s)
}
