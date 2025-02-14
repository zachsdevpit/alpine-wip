package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/zachsdevpit/alpine-wip/internal/middleware"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	tmpl := template.Must(template.New("").ParseGlob("templates/*.html"))

	r := http.NewServeMux()

	resources := http.FileServer(http.Dir("./resources"))
	r.Handle("GET /resources/", http.StripPrefix("/resources", resources))

	r.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "home.html", nil)
	})

	r.HandleFunc("/menu", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "home.html", nil)
	})

	srv := &http.Server{
		Handler: middleware.Logging(logger)(r),
		Addr:    ":8080",
	}

	logger.Info("server listening", slog.String("addr", srv.Addr))

	srv.ListenAndServe()
}
