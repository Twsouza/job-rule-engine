package main

import (
	"fmt"
	"os"

	"github.com/Twsouza/job-rule-engine/application/router"
	"github.com/joho/godotenv"
)

var (
	port string
)

func init() {
	godotenv.Load()

	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
}

func main() {
	routes := router.SetupRouter()
	fmt.Printf("Server running on port %s\n", port)
	routes.Run(":" + port)
}
