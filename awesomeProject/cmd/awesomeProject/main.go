package main

import (
	"awesomeProject/internal/api"
	"log"
)

func main() {
	log.Println("Application startup")
	api.StartServer()
	log.Println("Application shutdown")
}
