package ch02

import (
	"net/http"
	"os"
)

func CallbackShutdown() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func shutdown(res http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}
