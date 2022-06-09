package main

import (
	"log"

	"github.com/eubyt/goeubyt/urlshort"
	"github.com/eubyt/goeubyt/web"
)

func startSQLite3() {
	urlshort.DATABASE.Connect("urlshort")
	urlshort.DATABASE.CreateTable()
}

func main() {
	startSQLite3()                  // Start SQLite3
	log.Panic(web.ListenAndServe()) // Iniciar servidor web
}
