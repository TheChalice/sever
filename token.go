package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	token      string
	tokenMutex sync.Mutex
)

func gettoken() string {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	return token

}

func updatatoken() {

	f := func() {
		v := url.Values{}
		v.Set("grant_type", "client_credential")
		v.Set("appid", "wx7597fee76b4d1d29")
		v.Set("secret", "1c5dc05d29ce7e13fc90f96ebf6bdbc4")
		//url:=url.URL
		r, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?" + v.Encode())
		if err != nil {
			return
		}
		if r != nil {
			defer r.Body.Close()
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}
		//GetResponseData(r)
		//json.Unmarshal
		var params = struct {
			Access  string `json:"access_token"`
			Expires int64  `json:"expires_in"`
		}{}
		err = json.Unmarshal(data, &params)
		if err != nil {
			return
		}

		log.Println("params", params)
		tokenMutex.Lock()
		token = params.Access
		tokenMutex.Unlock()

	}
	f()
	for range time.Tick(time.Hour) {
		f()
	}
}
