package main

import "GohCMS2/main/server"

func main() {
	router := server.InitServer()
	err := server.StartServer(router)
	if err != nil {
		panic(err)
	}
}
