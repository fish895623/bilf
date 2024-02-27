package route

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
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
	Convey("Test Setup", t, func() {
		Convey("Get Status Code", func() {
			So(w.Result().StatusCode, ShouldEqual, 200)
		})
		Convey("Get Test Value", func() {
			So(d.LowerCase, ShouldEqual, "asdf")
		})
	})
}
func TestPingGet(t *testing.T) {
	r := Setup()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	var body struct {
		Status     string  `json:"status"`
		DivPercent float32 `json:"DividendsPercentage"`
	}

	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Error(err)
	}

	Convey("Test Ping Get", t, func() {
		Convey("Get Status Code", func() {
			So(w.Result().StatusCode, ShouldEqual, 200)
		})
		Convey("Get Test Value", func() {
			So(body.DivPercent > 0, ShouldEqual, true)
		})
	})
}
func TestPingPost(t *testing.T) {
	var send_body struct {
		Name string `json:"name"`
	}

	send_body.Name = "JEPI"

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(send_body)

	r := Setup()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/ping", &buf)
	r.ServeHTTP(w, req)
}

func Test_nil_ping_post(t *testing.T) {
	var send_body struct {
		Name string `json:"name"`
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(send_body)

	r := Setup()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/ping", &buf)
	r.ServeHTTP(w, req)
}

func TestRequestingQuote(t *testing.T) {
	res := RequestingQuote("JEPI")
	if res.StatusCode != 200 {
		t.Error("Fail to get html")
	}
}
