package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	logs "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logs.Infof("Hello Test Echo Server")
		json.NewEncoder(w).Encode("Hello Test Echo Server")
	})

	router.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		logs.Infof("Hello Test Echo Server v2.0 @oversea")
		json.NewEncoder(w).Encode("Hello Test Echo Server v2.0 @oversea")
	})

	http.ListenAndServe(":8080", router)
}
