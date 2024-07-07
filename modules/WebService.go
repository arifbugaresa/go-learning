package modules

import (
	"fmt"
	"go-learning/middleware"
	"go-learning/modules/movies"
	"log"
	"net/http"
)

func LearningWebService() {
	WebService()
	WebServiceWithMiddleware()
}

func WebService() {
	http.HandleFunc("/movies", movies.GetMovies)
	http.HandleFunc("/post_movie", movies.PostMovie)

	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func WebServiceWithMiddleware() {
	// server configuration
	server := &http.Server{
		Addr: ":8080",
	}

	// routing
	http.Handle("/", middleware.Log(http.HandlerFunc(movies.GetMovies)))

	// run the server
	fmt.Println("server running at http://localhost:8080")

	server.ListenAndServe()
}
