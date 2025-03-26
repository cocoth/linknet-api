package main

import (
	"github.com/cocoth/linknet-api/src/utils"
)

func main() {
	a := utils.SanitizeString("HRA/AH/Sales_Central_Java/234/084 03-jan ")
	utils.Debug(a)
}
