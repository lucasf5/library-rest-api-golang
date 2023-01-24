package utils

import (
	"library-api-rest/src/database"
	"library-api-rest/src/model"
)

var BookModel model.Book

func CreatBookMock() {
	aluno := model.Book{
		Title:            "Teste",
		Isbn:             "123456789",
		PageCount:        123,
		Status:           "PUBLISH",
		ShortDescription: "Teste",
		LongDescription:  "Teste",
	}

	database.Connection().Create(&aluno)
	BookModel = aluno
}

func DeleteBookMock() {
	database.Connection().Delete(&BookModel, BookModel.Id)
}
