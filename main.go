package main

import (
	"fmt"
	"client"
)

func main() {
	client.Init()
	fmt.Println(client.Query("google.com"))
	fmt.Println(client.Query("quora.com"))
	fmt.Println(client.Query("fb.com"))
}
