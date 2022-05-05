package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Printf(err.Error())
	}
}
