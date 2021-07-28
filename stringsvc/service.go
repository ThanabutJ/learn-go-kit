package main

import (
	"errors"
	"strings"
)

//StringService
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

//stringService
type stringService struct {
}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

var ErrEmpty = errors.New("Empty string")
