package main

import (
	"common-tools/config"
	"common-tools/logger"
	"common-tools/zflag"
	"fmt"
)

func main() {
	zflag.Parse()
	fmt.Println(config.Init(zflag.ConfigFile))
	fmt.Println(config.GlobalConfig)
	fmt.Println(logger.Init(config.GlobalConfig.Log.Level))
}
