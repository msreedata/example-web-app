// Copyright 2017 Ad Hoc

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			home(w, r)
		case "/_healthcheck":
			healthcheck(w, r)
		case "/envdump":
			envdump(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	log.Printf("Listening on http://0.0.0.0:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()
	tmpl, err := template.New("hello-world.html").Parse(`
<!doctype html>
<html lang=en>
    <head>
        <meta charset=utf-8>
        <title>Hello, {{ .Name }}</title>
        <style>
          html, body {
              background: hsla(155, 24%, 78%, 0.9);
              color: rgba(9, 87, 54, 0.9);
              font-family: "Segoe UI", Frutiger, "Frutiger Linotype", "Dejavu Sans", "Helvetica Neue", Arial, sans-serif;
          }
          body {
              width: 800px;
              margin: 0 auto;
              padding: 40px 0;
          }
h1 {
    font-family: "Segoe UI", Frutiger, "Frutiger Linotype", "Dejavu Sans", "Helvetica Neue", Arial, sans-serif;
    font-weight: 400;
    font-size: 48px;
    color: rgba(52, 43, 24, 0.9);
    background: hsla(119, 68%, 96%, 0.9);
    padding: 10px;
    text-align: center;
}
        </style>
    </head>
    <body>
        <h1>Hello, {{ .Name }}</h1>
	<hr>
	<p>Hello, This is a Go hello world page!</p>
	<p>Try setting the name with a URL parameter <code>?name=FOO</code>.</p>
        <p>Examples:</p>
        <ul>
            <li><a href=".">Hello, Go (default)</a>
            <li><a href="?name=docker%Studio">?name=docker Studio</a>
        </ul>
        <hr>
        <p><a href="/envdump">Dump environment variables</a></p>
        <hr>
        This page generated in {{ .Elapsed }} seconds.
    </body>
</html>
`)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	name := r.FormValue("name")
	if name == "" {
		name = "World"
	}
	if err := tmpl.Execute(w, struct {
		Name    string
		Elapsed time.Duration
	}{
		Name:    name,
		Elapsed: time.Now().Sub(t0),
	}); err != nil {
		log.Printf("executing template: %v", err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("OK\n"))
}

func envdump(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	for _, pair := range os.Environ() {
		fmt.Fprintf(w, "%s\n", pair)
	}
}
