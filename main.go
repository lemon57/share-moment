package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lemon57/share-moment/controllers"
	"github.com/lemon57/share-moment/templates"
	"github.com/lemon57/share-moment/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS, "tailwind.gohtml", "navbar.gohtml", "home.gohtml",
	))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS, "tailwind.gohtml", "navbar.gohtml", "contact.gohtml",
	))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(
		templates.FS, "tailwind.gohtml", "navbar.gohtml", "faq.gohtml",
	))))

	r.Get("/signup", controllers.FAQ(views.Must(views.ParseFS(
		templates.FS, "tailwind.gohtml", "navbar.gohtml", "signup.gohtml",
	))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
