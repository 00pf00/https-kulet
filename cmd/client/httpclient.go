package main

import "00pf00/https-kulet/pkg/https/client"

func main() {
	client := client.NewClient()
	client.Post()
}
