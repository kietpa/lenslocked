package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/kiet-asmara/lenslocked/views"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	t, err := views.Parse(filepath)
	if err != nil {
		log.Printf("Parsing template %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("Templates", "home.gohtml")
	executeTemplate(w, tplPath) // generally make execute last action
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("Templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("Templates", "faq.gohtml")
	executeTemplate(w, tplPath)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("start at :3000")
	http.ListenAndServe(":3000", r)
}

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	case "/faq":
// 		faqHandler(w, r)
// 	default:
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	}
// }

// func main() {
// 	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("hello"))
// 	}

// 	http.HandleFunc("/", handlerIndex)
// 	address := "localhost:9000"
// 	fmt.Println("start server at ", address)
// 	err := http.ListenAndServe(address, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// }
