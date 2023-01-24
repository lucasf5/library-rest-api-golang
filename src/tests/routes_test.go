package tests

import (
	"encoding/json"
	"library-api-rest/src/controllers"
	"library-api-rest/src/database"
	"library-api-rest/src/model"
	"library-api-rest/src/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Request() *gin.Engine {
	r := HandleRequest()
	return r
}

func TestGetAll(t *testing.T) {
	r := Request()
	r.GET("/books", controllers.GetAll)
	req, _ := http.NewRequest("GET", "/books", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var books model.AllBooksResponse

	json.Unmarshal(response.Body.Bytes(), &books)

	assert.Equal(t, books.Embedded.Message, "Livros encontrados com sucesso.", "Não é igual a Livros encontrados com sucesso.")

}

func TestGetOne(t *testing.T) {
	r := Request()
	r.GET("/books/:id", controllers.GetOne)
	req, _ := http.NewRequest("GET", "/books/1", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var book model.BookResponse

	json.Unmarshal(response.Body.Bytes(), &book)

	assert.Equal(t, book.Embedded.Message, "Livro encontrado com sucesso.", "Não é igual a Livro encontrado com sucesso.")
	assert.Equal(t, book.Embedded.Book.Title, "Unlocking Android", "Não é igual a Unlocking Android")

}

func TestCreate(t *testing.T) {
	utils.CreatBookMock()

	database.Connection().Find(&utils.BookModel.Id, &utils.BookModel)

	assert.Equal(t, utils.BookModel.Title, "Teste", "Não é igual a Teste")

	utils.DeleteBookMock()
}

func TestUpdate(t *testing.T) {
	utils.CreatBookMock()

	database.Connection().Find(&utils.BookModel.Id, &utils.BookModel)

	newBook := &model.Book{
		Title:            "Teste 2",
		Isbn:             "123456789",
		PageCount:        123,
		Status:           "PUBLISH",
		ShortDescription: "Teste",
		LongDescription:  "Teste",
	}

	database.Connection().Save(&newBook)

	assert.Equal(t, newBook.Title, "Teste 2", "Não é igual a Teste 2")

	utils.DeleteBookMock()
}

func TestDelete(t *testing.T) {
	utils.CreatBookMock()

	utils.DeleteBookMock()

	assert.Equal(t, utils.BookModel.Id, int(0), "Não é igual a 0")

}

func TestFindByTitle(t *testing.T) {
	r := Request()
	r.GET("/books/search/:title", controllers.FindByTitle)
	req, _ := http.NewRequest("GET", "/books/search/Unlocking", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var books model.BookResponse

	json.Unmarshal(response.Body.Bytes(), &books)

	assert.Equal(t, books.Embedded.Book.Title, "Unlocking Android", "Não é Unlocking Android")

}
