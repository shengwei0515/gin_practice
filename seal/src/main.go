package main

import (
	"seal/routers"
)

func main() {

	router := routers.InitRouters()
	router.Run(":8080")
}
