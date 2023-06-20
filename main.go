package main

import (
	"fmt"
	"github.com/JackyBaooo/common-tools/zflag"
)

func main() {
	zflag.Parse()
	fmt.Println(zflag.ConfigFile)
}
