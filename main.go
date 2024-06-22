package main

import (
	"karma_auth/routes"
	// "karma_auth/config"
)

func main() {
	app := routes.Users()
	app.Listen(":3000")
}