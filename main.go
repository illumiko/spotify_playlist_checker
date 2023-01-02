package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func hand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<h1>HI</h1>`)
}

const base_url = "https://api.spotify.com/v1"
const client_id = "d225133173b2440db06090428ac2dea2"
const client_secret = "342f22ddec334217aa893d61ff6a322f"

func main() {
	http.HandleFunc("/", authorize)
	http.ListenAndServe(":8080", nil)

}

// {{{
func authorize(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	req, _ := http.NewRequest("GET", base_url+"/authorize", nil)
	req.Header.Set("client_id", client_id)
	req.Header.Set("response_type", "code")
	req.Header.Set("redirect_uri", "http://localhost:8080/")
	client.Do(req)
	http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)

}

func request_spotify(url string) *http.Response {
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", client_id)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
} // }}}
