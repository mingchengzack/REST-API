package main

import "github.com/mingchengzack/REST-API/api"

func main() {
	api := &api.API{}
	api.Init()
	api.Run(":8080")
}
