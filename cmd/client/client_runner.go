package main

import (
	"dns"
	"fmt"
)

func main() {
	client := dns.NewClient()
	result := client.Query("google.com")
	fmt.Println(result)
}
