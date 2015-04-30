package main

import (
	"data"
	"io"
	"net/http"
	"strings"
	"ui"
)

var d data.Data

func viewHandler(w http.ResponseWriter, r *http.Request) {
	requestUrl := r.URL.Path
	c := ui.Route(requestUrl, &d)
	responseHtml := c.RenderUiConfig(d)
	io.WriteString(w, string(responseHtml))
}

func main() {
	d = data.CreateData()
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "img/favicon.ico") })
	http.HandleFunc("/polymer-test/", func(w http.ResponseWriter, r *http.Request) {
		url := strings.Trim(r.URL.Path, "/")
		http.ServeFile(w, r, url)
	})
	http.HandleFunc("/polymer-test2/", func(w http.ResponseWriter, r *http.Request) {
		url := strings.Trim(r.URL.Path, "/")
		http.ServeFile(w, r, url)
	})
	http.HandleFunc("/", viewHandler)

	http.ListenAndServe(":8080", nil)
}
