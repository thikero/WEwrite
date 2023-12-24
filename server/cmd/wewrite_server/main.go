package main

import (
	"fmt"

	listener "github.com/thikero/WEwrite/internal/listener"
)

func main() {
	fmt.Println("Hello Goers")
	res := listener.Listen()
	fmt.Printf("Listener : %d", res)
}
