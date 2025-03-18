package main

import (
	"fmt"

	"github.com/cocoth/linknet-api/src/utils"
)

type user struct {
	Name string
}

func main() {
	var users user
	users.Name = "Coco"

	addr := &users.Name
	*addr = "Coco1"
	*addr = "Coco2"
	// addr.Name = "Coco2"

	fmt.Println(addr)
	utils.Info(users.Name)
}
