package main

import (
	"fmt"
	"github.com/summer-boythink/gokit/kithttp"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan struct{}, 1)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		// it is a mistake
		c <- struct{}{}
		kithttp.SimplePage(7777)
	}()

	for {
		select {
		case <-sig:
			log.Println("process exit...")
			os.Exit(1)
		case <-c:
			res := kithttp.Get("http://localhost:7777")
			fmt.Print(res)
			res = kithttp.PostWithJson("http://localhost:7777", "")
			fmt.Print(res)
			return
		}
	}
}
