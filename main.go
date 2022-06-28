package main

import (
	"fmt"

	"github.com/DianaCandy/FileProvider/config"
)

func main() {
	fmt.Println("Reading config")
	config.Init()
}
