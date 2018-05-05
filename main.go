package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const addr = ":12052"

const configFile = "/etc/shadowsocks.json"

//const configFile = "/Users/fuxiaosong/develop/go/gopath/src/ssinfo/ssdocker/shadowsocks_single_sample.json"

func main() {
	log.Printf("ssinfo server start success, and listening on localhost%s ...", addr)

	http.HandleFunc("/ssinfo", ssInfo)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func ssInfo(writer http.ResponseWriter, http *http.Request) {
	file, err := os.Open(configFile)
	if err != nil {
		writer.Write([]byte(err.Error()))
		return
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		writer.Write([]byte(err.Error()))
		return
	}

	writer.Write(bytes)
}
