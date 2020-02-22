package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mux "net/http"

	"github.com/stretchr/testify/assert"
)

var router *mux.ServeMux

func init() {

	router = CreateRoutes()
}

func TestGreetings(t *testing.T) {

	// When: GET /greet is called
	req, _ := http.NewRequest("GET", "/greet", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Then: status is 200
	assert.Equal(t, http.StatusOK, rr.Code)

	// And: Body contains greetings
	expected := `{ text: 'Greetings from docker container'}`
	assert.Equal(t, expected, rr.Body.String())
}
