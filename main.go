package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load env files", err.Error())
	}
}

func main() {
	fmt.Println("Run serveur")
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Exec route hello")
	fmt.Fprintf(w, os.Getenv("HELLO"))
}
