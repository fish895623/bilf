package route

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func testSetup() (r *gin.Engine, w *httptest.ResponseRecorder, req *http.Request) {
	r = Setup()

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)
	return
}

func TestSetup(t *testing.T) {
	_, w, _ := testSetup()
	var d struct {
		LowerCase string `json:"asdf"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &d)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, d.LowerCase, "asdf")
}
