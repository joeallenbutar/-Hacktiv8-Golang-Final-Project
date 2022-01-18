//Joe Allen Butarbutar (GLNG020ONL003)

package main

import (
	"Final-Project/database"
	"Final-Project/routers"
	"log"
)

// @title           Swagger ToDo's API
// @version         1.0
// @description     This is a sample server create todo's application.
// @termsOfService  http://swagger.io/terms/
// @contact.name   	API Support
// @contact.url   	http://www.swagger.io/support
// @contact.email 	support@swagger.io
// @license.name  	Apache 2.0
// @license.url   	http://www.apache.org/licenses/LICENSE-2.0.html
// @host      		localhost:8080
// @BasePath  		/

func main() {
	database.Setup()
	r := routers.Setup()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
