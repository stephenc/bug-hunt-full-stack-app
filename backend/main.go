package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	var err error
	db, err = initDB()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/todos", todosHandler)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
