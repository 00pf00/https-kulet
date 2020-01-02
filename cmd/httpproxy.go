package main

import (
	"00pf00/https-kulet/pkg/https/client"
)

func main() {
	httpclient := client.NewClient()
	httpclient.Post()
}
