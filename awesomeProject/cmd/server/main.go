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

// @title Прогнозы
// @version 1.0
// @description Сервис прогнозирования погодных параметров (условий)

// @contact.name Рабраб
// @contact.url https://github.com/NikolayB800H
// @contact.email gorkunovnm@gmail.com

// @license.name AS IS (NO WARRANTY)

// @host 0.0.0.0:8084
// @schemes https http
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app, err := app.New()
	if err != nil {
		log.Println("app can not be created", err)
		return
	}
	app.Run()
}
