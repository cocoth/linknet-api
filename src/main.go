package main

import (
	"github.com/cocoth/linknet-api/cmd"
	"github.com/cocoth/linknet-api/src/database"
	"github.com/cocoth/linknet-api/src/utils"
)

func init() {
	utils.LoadEnv()
	database.ConnectToDB()
}

func main() {
	cmd.Exec()
}
