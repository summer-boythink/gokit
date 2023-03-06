package kithttp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PostMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// SimplePage provides a hello world page(GET/POST) on the specified port,
// default port 8090
func SimplePage(port ...int) {
	if len(port) > 1 {
		log.Fatal("port only a param")
	}
	if len(port) == 0 {
		port = append(port, 8090)
	}
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			io.WriteString(w, "Hello, world!\n")
		} else if req.Method == "POST" {
			body, err := io.ReadAll(req.Body)
			if err != nil {
				log.Fatalln(err)
			}
			postResult := PostMessage{200, "success", string(body)}
			res, err := json.Marshal(postResult)
			if err != nil {
				log.Fatalln(err)
			}
			res = append(res, '\n')
			w.Header().Set("Content-Type", "application/json")
			io.WriteString(w, string(res))
		} else {
			io.WriteString(w, "No Support Method\n")
		}

	}

	http.HandleFunc("/", helloHandler)
	log.Printf("http start in,please visit http://localhost:%d", port[0])
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port[0]), nil))
}
