package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "../templates/*"))

func EnvironmentCheck() {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current Directory:", dir)

	// Read directory entries
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// List files and directories
	for _, entry := range entries {
		info := ""
		if entry.IsDir() {
			info = "[DIR] "
		} else {
			// Optionally get file info
			fileInfo, err := entry.Info()
			if err == nil {
				info = fmt.Sprintf("[FILE] (%d bytes) ", fileInfo.Size())
			}
		}
		fmt.Println(info + entry.Name())
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}

		err := t.ExecuteTemplate(w, "index.html.tmpl", data)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
			return
		}
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
