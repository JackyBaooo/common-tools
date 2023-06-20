package zflag

import (
	"flag"
	"os"
)

var (
	// internal variables
	help bool
	// external variables
	ConfigFile string
)

func Parse() {
	flag.BoolVar(&help, "h", false, "show help")
	flag.StringVar(&ConfigFile, "c", "config.json", "config file")
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}
}
