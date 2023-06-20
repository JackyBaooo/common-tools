package main

import (
	"fmt"
	"github.com/JackyBaooo/common-tools/config"
	"github.com/JackyBaooo/common-tools/logger"
	"github.com/JackyBaooo/common-tools/zflag"
)

func main() {
	zflag.Parse()
	fmt.Println(config.Init(zflag.ConfigFile))
	fmt.Println(config.GlobalConfig)
	fmt.Println(logger.Init(config.GlobalConfig.Log.Level))
}
