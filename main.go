package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		id := strings.TrimPrefix(r.URL.Path, "/articles/")
		if len(id) == 0 {
			json.NewEncoder(w).Encode(Articles)
		} else if len(id) >= 1 {
			for i, article := range Articles {
				if article.Id == id && i > -1 {
					json.NewEncoder(w).Encode(article)
				}
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	case "POST":
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		id := r.FormValue("id")
		title := r.FormValue("title")
		desc := r.FormValue("desc")
		cont := r.FormValue("content")
		Articles = append(Articles, Article{Id: id, Title: title, Desc: desc, Content: cont})
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called "}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func querryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		q := r.FormValue("q")
		for i, article := range Articles {
			if strings.Contains(article.Title, q) && i > -1 {
				json.NewEncoder(w).Encode(article)
			} else if strings.Contains(article.Desc, q) {
				json.NewEncoder(w).Encode(article)
			} else if strings.Contains(article.Content, q) {
				json.NewEncoder(w).Encode(article)
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}


func handleRequests() {
	http.HandleFunc("/articles/", handler)
	http.HandleFunc("/articles/search/", querryHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

	Articles = []Article{
		Article{Id: "1", Title: "Hello adarsh", Desc: "Article Description super", Content: "Article Content major"},
		Article{Id: "2", Title: "Hello akshay", Desc: "Article Description good", Content: "Article Content minor"},
	}

	handleRequests()
}
