package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ReadJson(jsonName string) (rawJson []byte, err error) {
	return os.ReadFile(jsonName)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path == "/favicon.ico" {
		return
	}
	splitPath := strings.Split(r.URL.Path, "/")[1:]
	// if len(splitPath) == 0 {
	// 	return
	// }
	rawJson, err := ReadJson(splitPath[0])
	if err != nil {
		fmt.Fprint(w, "Error json fille no exist : ", splitPath[0])
		return
	}

	if len(splitPath) == 1 {
		fmt.Fprint(w, string(rawJson))
		return
	}
	var data interface{}
	var newDataSend interface{}
	err = json.Unmarshal(rawJson, &data)
	if err != nil {
		fmt.Fprint(w, "Error json fille parse err : ", err)
		return
	}
	fmt.Println(data)

	for i := 1; i < len(splitPath); i++ {
		newData, ok := data.(map[string]interface{})[splitPath[i]]
		if !ok {
			fmt.Fprint(w, "Error proprety ", splitPath[0], ".", strings.Join(splitPath[1:i+1], "."), " in fille ", splitPath[0], " no exist.")
			return
		}
		newDataSend = newData
	}
	fmt.Fprint(w, (newDataSend))
}
