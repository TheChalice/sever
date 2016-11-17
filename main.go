package main

import (
	"crypto/sha1"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}
func weixinin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	log.Println(r.Form) //这些信息是输出到服务器端的打印信息
	log.Println("path", r.URL.Path)
	log.Println("scheme", r.URL.Scheme)
	log.Println(r.Form["url_long"])
	// winxin
	// 5oF8FPAe3Z9kNmnqwq1EKfrI58yvC1gUd6KZBAKDkdr

	for k, v := range r.Form {
		log.Println("key:", k)
		log.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprint(w, "hello sb ", checkSignature(r)) //这个写入到w的是输出到客户端的
}
func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/interface", weixinin)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func checkSignature(r *http.Request) bool {
	signature := r.FormValue("signature")
	timestamp := r.FormValue("timestamp")
	nonce := r.FormValue("nonce")

	token := "winxin"
	tmpArr := sort.StringSlice{token, timestamp, nonce}
	sort.Sort(tmpArr)
	tmpStr := strings.Join(tmpArr, "")

	//s := "sha1 this string"
	//产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。
	h := sha1.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(tmpStr))
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	//SHA1 值经常以 16 进制输出，例如在 git commit 中。使用%x 来将散列结果格式化为 16 进制字符串。
	fmt.Println(string(bs))
	fmt.Printf("%x\n", bs)
	if string(bs) == signature {

		return true
	} else {
		return false
	}
}
