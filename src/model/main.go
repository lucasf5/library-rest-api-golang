package model

import (
	"gopkg.in/validator.v2"
)

type Main struct {
	Links struct {
		Self string `json:"self"`
	} `json:"_links"`
	Message string `json:"message"`
}

type AllBooksResponse struct {
	Embedded struct {
		Message string `json:"message"`
		Books   []Book `json:"books"`
	} `json:"embedded"`
	Links struct {
		Self string `json:"self"`
	} `json:"_links"`
}

type BookResponse struct {
	Embedded struct {
		Message string `json:"message"`
		Book    Book   `json:"book"`
	} `json:"embedded"`
	Links struct {
		Self string `json:"self"`
	} `json:"_links"`
}

type Book struct {
	Id               int    `json:"id" validate:"nonzero"`
	Title            string `json:"title" validate:"nonzero,required"`
	Isbn             string `json:"isbn" validate:"nonzero"`
	PageCount        int    `json:"pageCount" validate:"nonzero"`
	Status           string `json:"status" validate:"nonzero"`
	ShortDescription string `json:"shortDescription" validate:"nonzero"`
	LongDescription  string `json:"longDescription" validate:"nonzero"`
}

type Books []Book

func Validate(book Book) error {
	if err := validator.Validate(book); err != nil {
		return err
	}

	return nil
}
