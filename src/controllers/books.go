package controllers

import (
	"library-api-rest/src/database"
	"library-api-rest/src/model"
	"net/http"

	"library-api-rest/src/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @host            localhost:8080
// @BasePath        /
// @schemes         http

// @Summary         Main
// @Description     Welcome to the API
// @Tags            Main
// @Accept          json
// @Produce         json
// @Success         200 {object} model.Main
// @Router          / [get]
func Main(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"_links": gin.H{
			"self": "http://localhost:8080",
		},
		"message": "Welcome",
	})
}

// @Summary         Get all books
// @Description     Get all books
// @Tags            Books
// @Accept          json
// @Produce         json
// @Success         200 {object} model.AllBooksResponse
// @Router          /books [get]

func GetAll(c *gin.Context) {
	var books model.Books
	database.Connection().Find(&books)
	if len(books) <= 0 {
		c.JSON(404, gin.H{"status": http.StatusNotFound, "message": "No books found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"_links": gin.H{
			"self": "http://localhost:8080/books",
		},
		"embedded": gin.H{
			"message": "Livros encontrados com sucesso.",
			"books":   books,
		},
	})
}

// @Summary         Get one book
// @Description     Get one book
// @Tags            Books
// @Accept          json
// @Produce         json
// @Param           id path int true "Book ID"
// @Success         200 {object} model.BookResponse
// @Router          /books/{id} [get]
func GetOne(c *gin.Context) {
	var book model.Book
	id, _ := c.Params.Get("id")

	database.Connection().Find(&book, id)

	if book.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Nenhum livro com id" + id + "foi encontrado",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"_links": gin.H{
			"self": "http://localhost:8080/books/" + id,
		},
		"embedded": gin.H{
			"message": "Livro encontrado com sucesso.",
			"book":    book,
		},
	})

}

// @Summary         Create a book
// @Description     Create a book
// @Tags            Books
// @Accept          json
// @Produce         json
// @Param           book body model.Book true "Book"
// @Success         201 {object} model.BookResponse
// @Router          /books [post]
func Create(c *gin.Context) {
	var book model.Book
	c.BindJSON(&book)
	if err := model.Validate(book); err != nil {
		utils.Erros(c, err)
		return
	}
	database.Connection().Create(&book)
	c.JSON(http.StatusCreated, gin.H{
		"_links": gin.H{
			"self": "http://localhost:8080/books/",
		},
		"embedded": gin.H{
			"message": "Livro criado com sucesso",
			"book":    book,
		},
	})
}

// @Summary         Update a book
// @Description     Update a book
// @Tags            Books
// @Accept          json
// @Produce         json
// @Param           id path int true "Book ID"
// @Param           book body model.Book true "Book"
// @Success         200 {object} model.BookResponse
// @Router          /books/{id} [put]
func Update(c *gin.Context) {
	var book model.Book
	id := c.Param("id")
	if err := model.Validate(book); err != nil {
		utils.Erros(c, err)
	}
	database.Connection().Find(&book, id)
	c.BindJSON(&book)
	database.Connection().Save(&book)

	c.JSON(http.StatusOK, gin.H{
		"_links": gin.H{
			"self": "http://localhost:8080/books/",
		},
		"embedded": gin.H{
			"message": "Livro atualizado com sucesso",
			"book":    book,
		},
	})
}

// @Summary         Delete a book
// @Description     Delete a book
// @Tags            Books
// @Accept          json
// @Produce         json
// @Param           id path int true "Book ID"
// @Success         200 {object} model.BookResponse
// @Router          /books/{id} [delete]
func Delete(c *gin.Context) {
	var book model.Book
	id := c.Param("id")
	database.Connection().Find(&book, id)
	if book.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Nenhum livro com id" + id + "foi encontrado",
		})
	}
	database.Connection().Delete(&book, id)
	c.JSON(http.StatusOK, gin.H{
		"_links": gin.H{
			"self": "http://localhost:8080/books/" + id,
		},
		"embedded": gin.H{
			"message": "Livro deletado com sucesso",
			"book":    book,
		},
	})
}

// @Summary         Search a book by title
// @Description     Search a book by title
// @Tags            Books
// @Accept          json
// @Produce         json
// @Param           title path string true "Book title"
// @Success         200 {object} model.BookResponse
// @Router          /books/search/{title} [get]
func FindByTitle(c *gin.Context) {
	var book model.Book
	title := c.Param("title")
	database.Connection().Where("title LIKE ?", "%"+title+"%").Find(&book)
	if book.Id == 0 {
		c.JSON(404, gin.H{"status": http.StatusNotFound, "message": "No books found!"})
		return
	}
	c.JSON(200, gin.H{
		"_links": gin.H{
			"self": "http://localhost:8080/books/search/" + title,
		},
		"embedded": gin.H{
			"message": "Livros encontrados com sucesso.",
			"book":    book,
		},
	})

}
