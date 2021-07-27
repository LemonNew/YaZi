package common

import (
	"../config"
	"flag"
)

func Banner() {
	banner := `
   ___                              _    
  / _ \     ___  ___ _ __ __ _  ___| | __ 
 / /_\/____/ __|/ __| '__/ _` + "`" + ` |/ __| |/ /
/ /_\\_____\__ \ (__| | | (_| | (__|   <    
\____/     |___/\___|_|  \__,_|\___|_|\_\   
                     fscan version: 1.6.3
`
	print(banner)
}

func Flag(Info *config.Info) {
	flag.StringVar(&Info.Url, "u", "", "http://127.0.0.1")
	flag.StringVar(&Info.Cms, "c", "", "PhpCms")
	flag.StringVar(&Info.ShellPassword, "pass", "pass", "Webshell pass")
	flag.Parse()
}
