package controller // import "github.com/FINTprosjektet/fint-tech-docs-service/controller"

import (
	"net/http"
	"log"
	"os"
)

func router(webroot string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request:", r.Method, r.URL, r.Proto)
		if r.Method == "POST" {
			if r.URL.Path == "/webhook" {
				GitHubWebHook(w,r)
			} else if r.URL.Path == "/api/projects/build" {
				BuildAllProjects(w,r)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else if r.Method == "GET" {
			if r.URL.Path == "/config.yml" {
				w.WriteHeader(http.StatusForbidden)
			} else if r.URL.Path == "/api/projects" {
				GetAllProjects(w,r)
			} else if _, err := os.Stat(webroot + r.URL.Path); err == nil {
				log.Println("Serving file", webroot + r.URL.Path)
				http.ServeFile(w, r, webroot + r.URL.Path)
			} else if _, err := os.Stat("." + r.URL.Path); err == nil {
				log.Println("Serving file", "." + r.URL.Path)
				http.ServeFile(w, r, "." + r.URL.Path)
			} else {
				http.ServeFile(w, r, webroot + "/index.html")
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

// SetupRouters ...
func SetupRouters(webroot string) http.Handler {
	log.Println("Setting up HTTP handler for webroot", webroot)
	return http.HandlerFunc(router(webroot))
}
