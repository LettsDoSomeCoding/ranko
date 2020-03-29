package main

import (
	"fmt"

	"github.com/LettsDoSomeCoding/ranko/api"
)

var greeting = "Hello I'm Ranko. I don't have much of a brain yet, but maybe I could remember a player or two?"

func main() {
	fmt.Println(greeting)
	api.StartServer(greeting)
}
