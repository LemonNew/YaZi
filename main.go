package main

import (
	"fmt"
	"github.com/loulan-ling/YaZi/common"
	"github.com/loulan-ling/YaZi/config"
)

func main() {
	var Info config.Info
	common.Flag(&Info)
	fmt.Println(Info.Url)
}
