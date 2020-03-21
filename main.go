package main

import (
	"github.com/mingchengzack/REST-API/api"
)

func main() {
	api := api.NewAPI()
	api.Run(":8080")
}
