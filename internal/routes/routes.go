package routes

import (
	"encoding/json"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/handlers/comments"
	"github.com/CVWO/sample-go-app/internal/handlers/posts"
	"github.com/CVWO/sample-go-app/internal/handlers/users"
	"github.com/go-chi/chi"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/users", GetUsers)
		r.Post("/users", NewUser)
		r.Get("/posts", GetPosts)
		r.Post("/posts", CreatePost)
		r.Delete("/posts/{id}", DeletePost)
		r.Get("/comments", GetComments)
		r.Post("/comments", CreateComment)
		r.Put("/comments/{id}", EditComment)
		r.Delete("/comments/{id}", DeleteComment)
	}
}

func GetUsers(w http.ResponseWriter, req *http.Request) {
	response, _ := users.HandleList(w, req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func NewUser(w http.ResponseWriter, req *http.Request) {
	users.HandleNewUser(w, req)
}

func GetPosts(w http.ResponseWriter, req *http.Request) {
	response, _ := posts.HandleList(w, req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreatePost(w http.ResponseWriter, req *http.Request) {
	posts.HandleCreatePost(w, req)
}

func DeletePost(w http.ResponseWriter, req *http.Request) {
	posts.HandleDeletePost(w, req)
}

func GetComments(w http.ResponseWriter, req *http.Request) {
	response, _ := comments.HandleList(w, req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateComment(w http.ResponseWriter, req *http.Request) {
	comments.HandleCreateComment(w, req)
}

func EditComment(w http.ResponseWriter, req *http.Request) {
	comments.HandleEditComment(w, req)
}

func DeleteComment(w http.ResponseWriter, req *http.Request) {
	comments.HandleDeleteComment(w, req)
}
