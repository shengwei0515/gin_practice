package main

import (
	"seal/router"
)

func main() {
	r := router.InitRouter()

	r.Run(":8080")
}
