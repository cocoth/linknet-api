package cmd

import (
	"fmt"
	"time"

	"github.com/cocoth/linknet-api/src/database"
	"github.com/spf13/cobra"
)

var dropdbCmd = &cobra.Command{
	Use:   "dropdb",
	Short: "Drop the database",
	Long:  "Drop the database",
	Run: func(cmd *cobra.Command, args []string) {
		DropDatabase()
	},
}

func init() {
	RootCmd.AddCommand(dropdbCmd)
}

func DropDatabase() {
	fmt.Print("Flushing Database...\n\n")
	time.Sleep(1 * time.Second)
	database.DropTables()
	database.ConnectToDB()
	fmt.Println("Database flushed successfully!")
	InitializeAndRunServer()
}
