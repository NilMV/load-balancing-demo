package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"
)

var counter int
var mutex = &sync.Mutex{}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	logrus.Info(strconv.Itoa(counter))
	mutex.Unlock()
}

func recieveCrashReport(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		incrementCounter(w, r)

	default:
		logrus.Info(w, "Sorry, only GET method is supported.")
	}
}

func main() {
	var webPort = flag.String("web_port", ":8080", "")

	http.HandleFunc("/", recieveCrashReport)

	logrus.Info(fmt.Sprintf("Starting simple server on :%v", *webPort))

	if err := http.ListenAndServe(*webPort, nil); err != nil {
		logrus.Fatalf("Failed to serve: %v", err)
	}
}
