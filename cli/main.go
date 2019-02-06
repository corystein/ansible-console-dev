package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
)

var publicAssets = packr.NewBox("../web/public/")

func main() {

	host := flag.String("h", "127.0.0.1", "port to serve on")
	port := flag.String("p", "3000", "port to serve on")
	//directory := flag.String("d", "./dist", "the directory of static file to host")
	flag.Parse()

	//http.Handle("/", http.FileServer(http.Dir(*directory)))

	http.Handle("/", http.FileServer(publicAssets))

	//log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)

	log.Printf("Serving on http://%s:%s\n", *host, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}
