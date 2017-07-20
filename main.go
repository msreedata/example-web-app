// Copyright 2017 Ad Hoc

package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t0 := time.Now()
		tmpl, err := template.New("hello-world.html").Parse(`
<!doctype html>
<html lang=en>
    <head>
        <meta charset=utf-8>
        <title>Hello, {{ .Name }}</title>
        <style>html, body { background: #116287; color: rgba(255, 255, 255, 0.9); font-family: -apple-system, BlinkMacSystemFont, sans-serif; }</style>
    </head>
    <body>
        <h1>Hello, {{ .Name }}</h1>
	<hr>
	<p>Try setting the name with a URL parameter <code>?name=FOO</code>.</p>
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
	})

	http.HandleFunc("/_healthcheck", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK\n"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
