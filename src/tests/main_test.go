package tests

import (
	"encoding/json"
	"library-api-rest/src/controllers"
	"library-api-rest/src/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func HandleRequest() *gin.Engine {
	r := gin.Default()
	return r
}

func TestMain(t *testing.T) {
	r := HandleRequest()
	r.GET("/", controllers.Main)
	req, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var main model.Main

	json.Unmarshal(response.Body.Bytes(), &main)

	assert.Equal(t, main.Message, "Welcome", "Não é igual a Welcome")
}
