package kithttp

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func PostWithJson(url string, jsonStr string) (res string) {
	buf := strings.NewReader(jsonStr)
	resp, err := http.Post(url, "application/json", buf)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	resp.Body.Close()
	return string(body)
}

func Get(url string) (res string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	resp.Body.Close()
	return string(body)
}
