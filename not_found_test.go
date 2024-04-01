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

func TestNotFound(t *testing.T) {

	router := httprouter.New()
	router.NotFound = (http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "gak ketemu")
	}))

	r := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, r)
	response := recoder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "gak ketemu", string(bytes))

}
