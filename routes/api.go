package routes

import (
	"net/http"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	panic("Not Implemented")
}

func ApiRouter(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/posts", GetPostsHandler)
}
