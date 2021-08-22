package main

import (
	"errors"
	"strings"
)

// StringService interface
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// ErrEmpty error
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for StringService
type ServiceMiddleware func(StringService) StringService
