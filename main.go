package main

import "fmt"

// "fmt"
// "server"

func main() {
	// Init()
	client := NewClient()
	result := client.Query("google.com")
	fmt.Println(result)

	// s := NewDnsServer()
	// s.Listen()
}
