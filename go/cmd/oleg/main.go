/*
Command oleg is a simple commandline interface to OlegDB.
*/
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Xe/oleg/go"
)

var (
	DB           *oleg.Database
	HostFlag     = flag.String("host", "127.0.0.1", "hostname olegdb is running on")
	PortFlag     = flag.String("port", "38080", "port olegdb is listening on")
	DeleteAction = flag.Bool("delete", false, "delete this key from olegdb?")
)

func main() {
	flag.Parse()

	DB = oleg.Purchase(*HostFlag, *PortFlag)

	var (
		table string
		key   string
		value string

		err error
	)

	table = flag.Arg(0)
	key = flag.Arg(1)

	if flag.NArg() == 3 {
		value = flag.Arg(2)
	} else if flag.NArg() == 0 {
		flag.Usage()
	}

	if value != "" {
		err = DB.Jar(table, key, value)
	} else {
		if *DeleteAction {
			err = DB.Scoop(table, key)
		} else {
			var val string
			val, err = DB.Unjar(table, key)
			if err == nil {
				fmt.Println(val)
			}
		}
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
