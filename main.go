package main

import (
	"go-facebook-bot/pkg/fb"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const projectDirName = "go-facebook-bot"

func getEnv() {
	err := godotenv.Load(`.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	//getEnv()
	http.HandleFunc("/webhook", fb.HandleMessenger)
	http.HandleFunc("/", homepageHandler)

	// port := ":4000"
	log.Println(os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":8089", nil))
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page here !!"))
}
