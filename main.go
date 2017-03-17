package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	indexTemplate := template.Must(template.ParseFiles("templates/index.tmpl"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		indexTemplate.Execute(w, map[string]interface{}{
			"imageURL": "https://pbs.twimg.com/media/C6r5dPxVoAEOil8.jpg",
		})
	}))
	srv := &http.Server{Addr: ":8081", Handler: mux}
	log.Printf("port: %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)
	}
}
