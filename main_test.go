package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"goBlog/config"

	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

var token string
var pwd = "12345678"
var emeil = "leiju@outlook.com"

var r *gin.Engine
var w *httptest.ResponseRecorder
func init() {
	r = setupRouter()
	w= httptest.NewRecorder()
}

func Test_PingRoute(t *testing.T) {

	
	req, _ := http.NewRequest("GET", `/v1/logn?pwd=`+pwd+`&emeil=`+emeil, nil)
	
	r.ServeHTTP(w, req)

	if assert.Equal(t, 200, w.Code) {
		token, _ = jsonparser.GetString(w.Body.Bytes(), "token")
	}
	code, _ := jsonparser.GetInt(w.Body.Bytes(), "code")

	assert.Equal(t, 200, int(code))

	req, _ = http.NewRequest("GET", `/blogs?page=1&per_page=5&author=`+emeil, nil)

	req.Header = map[string][]string{
		"Authorization": {"Bearer " + token},
	}
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if assert.Equal(t, 200, w.Code) {
		//assert.Error(t, nil, w.Body.String())
		code, _ := jsonparser.GetInt(w.Body.Bytes(), "code")
		assert.Equal(t, 200, int(code))
		//content,_, _, _ := jsonparser.Get(, "data")
		assert.Error(t, nil, w.Body.String())
	}
}

func Test_config(t *testing.T) {
	config.Set("test.leiju", "test")
}

func TestPingRoute(t *testing.T) {

	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
