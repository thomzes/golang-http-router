package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {

	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		text := "Product " + id
		fmt.Fprint(w, text)
	})

	r := httptest.NewRequest("GET", "http://localhost:3000/products/1", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, r)
	response := recoder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1", string(bytes))

}
