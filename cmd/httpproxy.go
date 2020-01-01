package main

import (
	"00pf00/https-kulet/pkg/https/client"
	"fmt"
	"os"
)

func main() {
	dir,_ := os.Getwd()
	fmt.Printf("workpath = %v",dir)
	httpclient := client.NewClient()
	httpclient.Post()
}
