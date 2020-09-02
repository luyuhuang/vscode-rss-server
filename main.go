package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var codeMap map[string]string
var mutex sync.Mutex

func setcode(w http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	state := req.URL.Query().Get("state")
	if len(code) == 0 || len(state) == 0 {
		fmt.Fprintf(w, "fail")
		return
	}

	mutex.Lock()
	codeMap[state] = code
	mutex.Unlock()

	log.Println("set", state, code)
	fmt.Fprintf(w, "success")
}

func getcode(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	state := req.URL.Query().Get("state")
	if len(state) == 0 {
		fmt.Fprintf(w, "{\"status\":-1}")
		return
	}

	mutex.Lock()
	code, ok := codeMap[state]
	if ok {
		delete(codeMap, state)
	}
	mutex.Unlock()

	if !ok {
		fmt.Fprintf(w, "{\"status\":1}")
	} else {
		log.Println("get", state, code)
		fmt.Fprintf(w, "{\"status\":0,\"code\":\"%s\"}", code)
	}
}

func main() {
	codeMap = make(map[string]string)

	http.HandleFunc("/setcode", setcode)
	http.HandleFunc("/getcode", getcode)

	addr := "127.0.0.1:8080"
	if len(os.Args) == 2 {
		addr = os.Args[1]
	}

	log.Println("listening at", addr)
	http.ListenAndServe(addr, nil)
}
