package main

import "github.com/alfiqo/disbursement/routes"

func main() {
	router := routes.NewResthandler()
	router.Listen(":8080")
}
