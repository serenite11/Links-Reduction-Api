package service

import (
	"errors"
	"net/url"
)

func validateLink(link string) error {
	if link == "" {
		return errors.New("Empty request")
	}
	if _, err := url.ParseRequestURI(link); err != nil {
		return errors.New("Url is not valid")
	}
	return nil
}
