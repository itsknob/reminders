package main

import (
	"fmt"
	"log"
	"testing"

	"net/http"

	"example.com/db"
	"example.com/logger"
	// "example.com/routes"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func fooHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Printf("GET - %s\n", path)
	fmt.Fprintf(w, "You've reached Foo")
	w.Write([]byte("Hello from /foo"))
}

func fooEchoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from /foo - echo")
}

type Reminder struct {
	Id, Name string
}

func main() {

	// DB Setup
	fmt.Println("Initializing Database")
	logger := &logger.Logger{}
	db, err := db.InitDb("test", logger)
	if err != nil {
		log.Fatal(err)
	}
	assert.NotNil(&testing.T{}, db, "DB Should be initialized")

	// Server Setup
	fmt.Printf("Starting Server")
	e := echo.New()
	e.GET("/foo", fooEchoHandler)
	// http.HandleFunc("/foo", fooHandler)

	// http.HandleFunc("/api", routes.ApiRouter)
	//
	// // last
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("GET - /")
	// 	fmt.Fprint(w, "Welcome to my website!")
	// })

	// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EsacpeString(r.URL.Path))
	// })

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Printf("Routes: %+v\n", http.DefaultServeMux)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
