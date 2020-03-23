package main

import (
	"fmt"

	"github.com/LettsDoSomeCoding/ranko/api"
)

func main() {
	fmt.Println("Hello I'm Ranko. I don't have a brain yet, but just you wait.")
	api.StartServer()
}

func log(x int, y int) {
	fmt.Println("hello")
}
