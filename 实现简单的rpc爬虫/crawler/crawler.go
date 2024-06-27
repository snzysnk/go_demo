package crawler

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Get(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	buf := make([]byte, 8*1024)
	for {
		n, err := resp.Body.Read(buf)
		if err == io.EOF {
			//read completed
			break
		}
		if err != nil {
			return result, err
		}
		result += string(buf[:n])
	}
	return result, err
}

func Storage(url string) {
	result, err := Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	hash := md5.New()
	hash.Write([]byte(url))
	fileName := hex.EncodeToString(hash.Sum(nil)) + ".html"
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	f.WriteString(result)
}

func Do(urls []string) {
	for _, url := range urls {
		go Storage(url)
	}
}
