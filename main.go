package main

import (
	"./central"
	"./common"
	"./config"
)

func main() {
	var Info config.Info
	common.Flag(&Info)
	central.Controller(&Info)
}
