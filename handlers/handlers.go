package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jegacs/simpleapi/models"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Response string `json:"response,omitempty"`
}

func SetHelloWorldHandler() {
	http.HandleFunc("/hello", HelloWorldHandler)
}

func SetShortenUrlHandler() {
	http.HandleFunc("/shortener", ShortenUrlHandler)
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received!")
	switch r.Method {
	case "GET":
		response := &Response{
			Response: "Hello, world",
		}
		serializedResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "There was an unexpected error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(serializedResponse)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

func ShortenUrlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received!")
	switch r.Method {
	case "POST":
		response := &Response{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "It seems your request is malformed", http.StatusBadRequest)
			return
		}

		payload := &models.CleanUriPayload{}
		err = json.Unmarshal(body, payload)
		if err != nil {
			http.Error(w, "It seems your request is malformed", http.StatusBadRequest)
			return
		}

		shortener := models.CleanUriAPI{}
		shortURL, err := shortener.Shorten(payload.URL)
		if err != nil {
			http.Error(w, "There was an unexpected error", http.StatusInternalServerError)
			return
		}

		if shortURL == "" {
			http.Error(w, "It seems you sent an invalid URL, please submit a valid one", http.StatusBadRequest)
			return
		}

		response.Response = shortURL
		serializedResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "There was an unexpected error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(serializedResponse)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

func Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, nil))
}