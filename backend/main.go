package main

import (
	"net/http"
	"log"
	"fmt"
)

func main()  {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponceWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}