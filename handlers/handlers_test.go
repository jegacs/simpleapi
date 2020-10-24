package handlers

import (
	"bytes"
	"github.com/jegacs/simpleapi/models"
	_ "github.com/smartystreets/assertions"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"encoding/json"
)

func TestHelloHandler(t *testing.T) {
	Convey("when doing a GET http method onto hello endpoint", t, func() {
		req, err := http.NewRequest("GET", "/hello", nil)
		So(err, ShouldBeNil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HelloWorldHandler)
		handler.ServeHTTP(rr, req)

		Convey("response should be status ok", func() {
			status := rr.Code
			So(status, ShouldEqual, http.StatusOK)
		})

		Convey("response message should be 'Hello world'", func() {
			body, err := ioutil.ReadAll(rr.Body)
			So(err, ShouldBeNil)
			response := &Response{}
			err = json.Unmarshal(body, response)
			So(err, ShouldBeNil)
			So(response.Response, ShouldEqual, "Hello, world")
		})
	})

	Convey("when doing methods other than GET onto hello endpoint", t, func() {
		Convey("status code should be 405 Method Not Allowed", func() {
			req, err := http.NewRequest("POST", "/hello", nil)
			So(err, ShouldBeNil)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HelloWorldHandler)
			handler.ServeHTTP(rr, req)
			So(rr.Code, ShouldEqual, http.StatusMethodNotAllowed)
		})
	})
}

func TestShortenUrlHandler(t *testing.T) {
	Convey("when doing POST http method onto /shorten endpoint ", t, func() {
		payload := &models.CleanUriPayload{
			URL: "https://www.google.com.mx",
		}

		serializedPayload, err := json.Marshal(payload)
		So(err, ShouldBeNil)

		req, err := http.NewRequest("POST", "/shorten", bytes.NewReader(serializedPayload))
		So(err, ShouldBeNil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ShortenUrlHandler)
		handler.ServeHTTP(rr, req)

		response := &Response{}
		serializedResponse, err := ioutil.ReadAll(rr.Body)
		So(err, ShouldBeNil)

		err = json.Unmarshal(serializedResponse, response)
		So(err, ShouldBeNil)

		So(response.Response, ShouldNotBeEmpty)

	})

	Convey("when doing http method other than POST onto /shorten endpoint", t, func() {
		Convey("status code should be 405 Method Not Allowed", func() {
			req, err := http.NewRequest("PUT", "/shorten", nil)
			So(err, ShouldBeNil)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(ShortenUrlHandler)
			handler.ServeHTTP(rr, req)
			So(rr.Code, ShouldEqual, http.StatusMethodNotAllowed)
		})
	})

	Convey("when doing POST with invalid body onto /shorten endpoint", t, func() {
		payload := &models.CleanUriPayload{
			URL: "",
		}

		serializedPayload, err := json.Marshal(payload)
		So(err, ShouldBeNil)

		req, err := http.NewRequest("POST", "/shorten", bytes.NewReader(serializedPayload))
		So(err, ShouldBeNil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ShortenUrlHandler)
		handler.ServeHTTP(rr, req)
		So(rr.Code, ShouldEqual, http.StatusBadRequest)
	})

}
