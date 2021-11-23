package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/album", getAlbum)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	io.WriteString(w, "hello")
}

type Album struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type St struct {
	id string
}

func (s *St) show() {
	println(s.id)
}

func getAlbum(w http.ResponseWriter, r *http.Request) {
	var st Album
	st.Id = "1"
	st.Name = "2"
	c, err := json.Marshal(st)
	if err == nil {
		_, err := io.WriteString(w, string(c))
		if err != nil {
			return
		}
	}
}
