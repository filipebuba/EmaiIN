package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		product := r.URL.Query().Get("product")
		id := r.URL.Query().Get("id")
		if product != "" {
			w.Write([]byte(product + id))
		} else {
			w.Write([]byte("Hello world"))
		}
	})
	r.Get("/{productName}/{productId}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "productName")
		w.Write([]byte(param))
	})
	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		obj := map[string]interface{}{"message": "sucess"}
		render.JSON(w, r, obj)

	})
	http.ListenAndServe(":3000", r)
}
