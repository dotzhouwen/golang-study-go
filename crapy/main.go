package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 Edg/99.0.1150.36"
)

type Ts struct {
	index string
	path string
}

// https://b1.szjal.cn/20201125/RgM9UXNf/1000kb/hls/OdUhvl3462001.ts

func main() {
	file := "D:\\video\\21\\index(3).m3u8"
	domain := "https://cdn3.jiuse.cloud/hls/609457/"
	local := "D:\\video\\21"

	flag.StringVar(&file, "file", "", "m3u8 location")
	flag.StringVar(&domain, "domain", "", "domain")
	flag.StringVar(&local, "local", "", "download dir")

	flag.Parse()

	if strings.TrimSpace(file) == "" || strings.TrimSpace(domain) == "" || strings.TrimSpace(local) == "" {
		return
	}

	tsFiles := readM3u8File(file, domain)

	var exitChan = make(chan bool)
	var tsFileChan = make(chan Ts, 3)

	go getPathFromChan(tsFileChan, exitChan, local)
	go sendPathToChan(tsFiles, tsFileChan)


	<-exitChan
}

func sendPathToChan(tsFiles []Ts, tsFileChan chan Ts) {
	for _, ts := range tsFiles {
		tsFileChan <- ts
	}
	close(tsFileChan)
}

func getPathFromChan(tsFileChan chan Ts, exitChan chan bool, basePath string) {
	for ch := range tsFileChan {
		fmt.Printf("download:%s\n", ch.path)
		sendRequest(ch.path, basePath, ch.index + ".ts")
	}
	exitChan <- true
}

func download(ts Ts) {
	fmt.Printf("download:%s -> %s\n", ts.index, ts.path)
	//time.Sleep(time.Second)
}

func sendRequest(url string, basePath string, downloadName string) {
	// "https://b1.szjal.cn/20201125/RgM9UXNf/1000kb/hls/OdUhvl3462001.ts"

	client := resty.New()
	req := client.R()
	req.SetHeader("user-agent", UserAgent)
	resp, err := req.Get(url)

	if err != nil {
		fmt.Println("error:", err)
	}

	downloadPath := filepath.Join(basePath, downloadName)
	if !FileExists(downloadPath) {
		err = ioutil.WriteFile(downloadPath, resp.Body(), 0666)
		if err != nil {
			return
		}
	}
}


func readM3u8File(filename string, domain string) []Ts {
	var tsPath []Ts
	file, err := os.Open(filename)
	if err != nil {
		log.Printf(err.Error())
	}
	reader := bufio.NewReader(file)

	var i = 0
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		readString = strings.TrimSpace(readString)
		if strings.HasSuffix(readString, ".ts") {
			index := fmt.Sprintf("%09d", i)
			ts := Ts{
				index: index,
				path: domain + readString,
			}
			tsPath = append(tsPath, ts)
			i++
		}
	}
	return tsPath
}

func FileExists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}