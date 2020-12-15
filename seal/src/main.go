package main

import (
	"seal/router"
	orm "seal/service/postgres"
)

func main() {
	defer orm.Db.Close()
	r := router.InitRouter()

	r.Run(":8080")
}
