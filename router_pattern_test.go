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

func TestRouterPatternNameParameter(t *testing.T) {

	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		itemId := p.ByName("itemId")
		text := "Product " + id + " Item " + itemId
		fmt.Fprint(w, text)
	})

	r := httptest.NewRequest("GET", "http://localhost:3000/products/1/items/1", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, r)
	response := recoder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1 Item 1", string(bytes))

}

func TestRouterPatternCatchAllParameter(t *testing.T) {

	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		image := p.ByName("image")
		text := "Image : " + image
		fmt.Fprint(w, text)
	})

	r := httptest.NewRequest("GET", "http://localhost:3000/images/small/profile.png", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, r)
	response := recoder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Image : /small/profile.png", string(bytes))

}
