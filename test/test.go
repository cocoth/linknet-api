package main

import "github.com/cocoth/linknet-api/src/utils"

func main() {
	errPass := utils.CompareHashPassword([]byte("11111111"), "$2a$10$iVXrf6oM3N5BSiWH6.dMmune6aB8qGcYW6TNoo4eiQ6dlr4AWj.L2db")
	if errPass != nil {

		utils.Debug(errPass.Error())
	}

	utils.Info("password matched")
}
