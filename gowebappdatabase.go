package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/amunnelly/gowebappdatabase/routing"
)

func main() {
	// TO-DO: Check about swapping prefixes
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", routing.ServeHome)

	// TO-DO: Fix these rascals
	fmt.Println("Incipio - I begin.")

	if  len(os.Getenv("PORT")) > 0 {
		log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
		} else {
			log.Fatal(http.ListenAndServe(":8080", nil))
		}
	}
