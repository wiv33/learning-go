package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)

	w.Header().Add("Set-Cookie", "VISIT=TRUE")

	if _, ok := r.Header["Cookie"]; ok {
		fmt.Fprintf(w, "<html><body>%s</body></html>", "두 번째 이후")
	} else {
		fmt.Fprintln(w, "<html><body>첫 번째</body></html>")
	}

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>%s</body></html>", "Hello World!")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening : 18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
