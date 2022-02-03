package webserver

import (
	"html/template"
	"log"
	"net/http"
)

func Webserver(port string) {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/", homePage)

	log.Println("Starting WEB server on port", port)
	err := http.ListenAndServe(":"+port, serverMux)

	if err != nil {
		log.Println("Starting WEB server failed:", err)
		return
	}

	defer func() {
		log.Println("Stopping WEB server on port", port)
	}()
}

// Show logs
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	renderTemplate(w, r, "webpage/index.html")
}

func renderTemplate(w http.ResponseWriter, r *http.Request, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println("Failed to serve page:", err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println("Failed to serve page:", err)
		return
	}
}
