package options

import (
	"flag"

	"github.com/cocoth/linknet-api/config"
)

func Opt() {
	dropDB := flag.Bool("drop-db", false, "Drop database and create new one")

	flag.Parse()

	if *dropDB {
		config.DropTables()
		config.ConnectToDB()
	}
}
