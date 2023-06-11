package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/healthz", Healthz)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
func Healthz(w http.ResponseWriter, r *http.Request) {

	ip := strings.Split(r.RemoteAddr, ":")[0]
	version := os.Getenv("VERSION")
	r_header := r.Header
	w.Header().Set("IP", ip)
	w.Header().Set("Version", version)
	for k, v := range r_header {
		w.Header().Set(k, v[0])

	}
	w.WriteHeader(200)
	response := "httpCode:200"
	w.Write([]byte(response))

}
