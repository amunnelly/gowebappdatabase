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
	fmt.Println(os.Getenv("PORT"))
	fmt.Println("Incipio - I begin.")

	// https://blog.gopheracademy.com/advent-2016/exposing-go-on-the-internet/

	if  len(os.Getenv("PORT")) > 0 {
		fmt.Println(os.Getenv("PORT"))
		log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
		} else {
			fmt.Println("8080, in theory.")
			log.Fatal(http.ListenAndServe(":8080", nil))
		}


		// log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

	// s := &http.Server{
	// 	Addr: ":8080",
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// 	IdleTimeout:  120 * time.Second,
	// }


	// if  len(os.Getenv("PORT")) > 0 {
	// 	s := &http.Server{
	// 		Addr: ":" + os.Getenv("PORT"),
	// 		ReadTimeout:  50 * time.Second,
	// 		WriteTimeout: 100 * time.Second,
	// 		IdleTimeout:  1200 * time.Second,
	// 	}
	// 	log.Fatal(s.ListenAndServe())
	// 		} 
			
	// 		log.Fatal(s.ListenAndServe())


}
