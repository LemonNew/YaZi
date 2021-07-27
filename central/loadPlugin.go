package central

import (
	"../config"


	. "../plugin/phpcms"

)




func load(Info *config.Info) {
	load_Phpcms(Info)
}

func load_Phpcms(Info *config.Info)  {
	Phpcms2008_type_getshell(Info)
}
