package main

import "net/http"

// Пример уместного использования функции init

func init() {
	redirect := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	http.HandleFunc("/blog", redirect)
	http.HandleFunc("/blog/", redirect)

	static := http.FileServer(http.Dir("static"))
	http.Handle("/favicon.ico", static)
	http.Handle("/fonts.css", static)
	http.Handle("/fonts/", static)
	http.Handle("/lib/godoc", http.StripPrefix(
		"/lib/godoc",
		http.HandlerFunc(staticHandler),
	))
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	// ...
}
