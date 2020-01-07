package main

import "00pf00/https-kulet/pkg/https/server"

func main() {
	server := server.NewHttpServer()
	server.StartServer()
}
