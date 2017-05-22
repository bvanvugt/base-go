package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type logConfig struct {
	Info  *log.Logger
	Debug *log.Logger
}

var logger = logConfig{}

func main() {
	logger.Info = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Lmicroseconds)
	logger.Debug = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)

	router := httprouter.New()
	router.GET("/", index)

	log.Fatal(http.ListenAndServe(":9000", router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
