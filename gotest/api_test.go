package gotest

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MakeAndRunWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "HELLO WORLD1")
	})
	return mux
}

/*
http test
*/
func TestWebHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	mux := MakeAndRunWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	body, _ := io.ReadAll(res.Body)
	assert.Equal(t, "HELLO WORLD", string(body))
}
