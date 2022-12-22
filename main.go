package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"qualifighting.backend.de/api/routes"
	"qualifighting.backend.de/lib"
)

func main() {
	router := routes.NewRouter()
	port := fmt.Sprintf(":%d", lib.GetAppConfig().Port)
	router.Run(port)
}
