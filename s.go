package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)

		var (
			queryString string
			referer     string
		)

		queryString = r.URL.RawQuery
		if queryString != "" {
			queryString = fmt.Sprintf("?%s", r.URL.RawQuery)
		}

		referer = r.Referer()
		if referer == "" {
			referer = "-"
		}

		msg := fmt.Sprintf(`%s "%s %s%s %s" "%s" %s`,
			r.RemoteAddr, r.Method, r.URL.Path, queryString, r.Proto, referer, r.UserAgent())
		log.Println(msg)
	})
}

func main() {

	var (
		host    string
		port    int
		docRoot string
	)
	flag.StringVar(&host, "h", "0.0.0.0", "listen ip or hostname")
	flag.IntVar(&port, "p", 8000, "listen port")
	flag.StringVar(&docRoot, "d", ".", "document root")
	flag.Parse()

	absDocRoot, err := filepath.Abs(docRoot)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Listening %s:%d on %s\n", host, port, absDocRoot)

	fs := http.FileServer(http.Dir(absDocRoot))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), loggingHandler(fs))
}
