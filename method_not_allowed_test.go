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

func TestMethodNotAllowed(t *testing.T) {

	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "gak boleh")
	})

	router.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "POST")
	})

	r := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, r)
	response := recoder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "gak boleh", string(bytes))

}
