package WebRequest

import (
	"../config"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func Request(Info *config.Info)(Result config.Result){


	//judge Url Is it empty
	if Info.Url != ""{
		results := VerifyUrl(Info.Url)
		if results != true{
			fmt.Println("Url is none")
			flag.Usage()
			os.Exit(0)
		}
	}
	if strings.Contains(Info.Url,"http://")==true{
		Result = Http(Info)
		return
	}else if strings.Contains(Info.Url,"https://")==true{
		Result = Https(Info)
		return
	}else {
		fmt.Println("Url is none")
		flag.Usage()
		os.Exit(0)
	}

	return
}

//Http Request
func Http(Info *config.Info)(Result config.Result){
	req, err := http.NewRequest("GET", Info.Url+Info.Payload,nil)
	req.Header.Set("User-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", Info.Cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("访问"+Info.Url+"失败")
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	Result.Body = string(body)
	Result.StatusCode = resp.StatusCode
	return Result
}


//Https Request
func Https(Info *config.Info)(Result config.Result){
	timeout := time.Duration(1000 * time.Second) //time out
	client  := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}


	req, err := http.NewRequest("GET", Info.Url+Info.Payload,nil)
	req.Header.Set("User-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", Info.Cookie)
	resp, err := client.Do(req)
	if err != nil {
			return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	Result.Body = string(body)
	Result.StatusCode = resp.StatusCode
	return Result


	return
}



//Verify url
func VerifyUrl (Url  string )(bool){
	urlRegex := "^(https?://)[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]"
	match, err := regexp.MatchString(urlRegex, Url)
	if err !=nil{
		return false
	}
	return match

}