package main

import (
	"net/http"

	"github.com/Alann07AS/basic_json_api/pkg/api"
	"github.com/Alann07AS/basic_json_api/pkg/args"
)

func main() {
	ip := args.ReadArgs(1)
	if ip == "" {
		ip = ":8080"
	}
	http.HandleFunc("/", api.MainHandler)
	http.ListenAndServe(ip, nil)
}
