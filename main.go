package main

import (
	"fmt"
	"log"
	"net/http"

	routes "github.com/aditya2305/REST_API/routers"
)

func main() {
	log.Fatal(http.ListenAndServe(":3000", routes.Router()))
	fmt.Printf("Listening on PORT 4000... \n")
}
