/*
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
*/
package main

import (
	app "awesomeProject/internal/pkg"
	"log"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Println("app can not be created", err)
		return
	}
	app.Run()
}
