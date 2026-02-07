package handlers

import (
	"fmt"
	"go1fl-sprint6-final/internal/service"
	"io"
	"net/http"
	"os"
	"time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Can't download the page", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/upload" {
		http.NotFound(w, r)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Can't download the form from the file", http.StatusInternalServerError)
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Can't read the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	result, err := service.ConvertMorse(string(data))
	if err != nil {
		http.Error(w, "Can't convert the file", http.StatusInternalServerError)
		return
	}

	nameFile := time.Now().UTC().Format("20060102150405") + ".txt"
	f, err := os.Create(nameFile)
	if err != nil {
		http.Error(w, "Can't create the file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = f.WriteString(result)
	if err != nil {
		http.Error(w, "Can't write the result in the file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, result)
}
