package materials

import (
	"errors"
	"publish/internal/publicate"
	"unicode/utf8"
)

var (
	errIdExist         = errors.New("ID already using")
	errNameEmpty       = errors.New("name must not be empty")
	errHeaderEmpty     = errors.New("header must not be empty")
	errPriceInvalid    = errors.New("price less than zero")
	ErrStatusPublished = errors.New("announcement already send")
	errEditDataEmpty   = errors.New("edit data must not be empty")
	ErrContextSmall    = errors.New("context in article must have more than 20 latters")
	ErrContentEmpty    = errors.New("description must not be empty")
)

func validateArticles(ID int, name, header string, articles map[int]publicate.Material) error {
	if _, ok := articles[ID]; ok {
		return errIdExist
	}

	if name == "" {
		return errNameEmpty
	}

	if header == "" {
		return errHeaderEmpty
	}

	return nil
}

func (a Article) validate(name, header, context string) error {
	if a.Status == "published" {
		return ErrStatusPublished
	}

	if name == "" && header == "" && context == "" {
		return errEditDataEmpty
	}

	return nil
}

func (a Announcement) validate(name, header, description string, price int) error {
	if a.Status == "published" {
		return ErrStatusPublished
	}

	if name == "" && header == "" && description == "" && price == -1 {
		return errEditDataEmpty
	}

	if price < -1 {
		return errPriceInvalid
	}

	return nil
}

func validateAnnouncement(name, header string, ID int, price int, a map[int]publicate.Material) error {
	if _, ok := a[ID]; ok {
		return errIdExist
	}

	if name == "" {
		return errNameEmpty
	}

	if price < 0 {
		return errPriceInvalid
	}

	if header == "" {
		return errHeaderEmpty
	}

	return nil
}

func (a Article) validatePublicate() error {
	if utf8.RuneCountInString(a.Context) < 20 {
		return ErrContextSmall
	}

	if a.Status == "published" {
		return ErrStatusPublished
	}

	return nil
}

func (a Announcement) validatePublicate() error {
	if a.Status == "published" {
		return ErrStatusPublished
	}

	if a.Content.Description == "" {
		return ErrContentEmpty
	}

	return nil
}

func validateAdd(message string) error {
	if message == "" {
		return errors.New("message is empty")
	}
	return nil
}
