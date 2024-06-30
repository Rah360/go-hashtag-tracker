package main

import (
	"encoding/json"
	"fmt"
	"hashtag_tracker/repository"
	"hashtag_tracker/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	repo := repository.NewInMemoryHashtagRepository()
	hashtagService := service.NewHashtagService(repo)
	r := mux.NewRouter()

	r.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		var postRequest struct {
			Content string `json:"content"`
		}
		if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		post := hashtagService.CreatePost(postRequest.Content)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)
	}).Methods("POST")

	r.HandleFunc("/hashtags/{hashtag}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hashtag := vars["hashtag"]
		count := hashtagService.GetHashtagCount(hashtag)
		response := map[string]int{hashtag: count}
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	r.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		posts := hashtagService.GetPosts()
		json.NewEncoder(w).Encode(posts)
	}).Methods("GET")

	// Start the server
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
