package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {

	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	r := httptest.NewRequest("GET", "http://localhost:3000/files/goodbye.txt", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, r)
	response := recoder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "goodbye httprouter", string(bytes))

}
