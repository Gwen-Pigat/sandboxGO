package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error fetch ENV")
	}
}

func main() {
	fmt.Println("Run serveur")
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Exec route hello")
	fmt.Fprint(w, os.Getenv("HELLO"))
}
