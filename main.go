package main

import (
	"fmt"
	"nlp_text_classifier/server"
)

func main() {
	fmt.Println("Boopsie!")
	str := server.Test("boopsie")
	fmt.Println(str)
	server.StartRouter()
}
