package main

import (
	"karma_auth/routes"
)

func main() {
	app := routes.Users()
	app.Listen(":3000")
}