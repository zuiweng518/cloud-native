package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func healthz(w http.ResponseWriter, r *http.Request) {
	host := r.RemoteAddr
	addr := strings.Split(host, ":")
	version := runtime.Version()
	// h := r.Header
	// for k, v := range h {
	// 	w.Header().Set(k, v[0])
	// }

	w.WriteHeader(200)
	w.Header().Add("Version:", version)
	fmt.Println("host:", addr[0])
	io.WriteString(w, "it work ok!")
}
