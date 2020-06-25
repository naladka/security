package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io"
	//"io/ioutil"
	"log"
	"net/http"
	"html/template"
	// "os/exec"
)

// Websocket message for gui
type Message struct {
	// Winner device number
	Message string `json:"Message"`
	// Device local addresses list
	DevicesLocAddr []string `json:"devices"`
	// Image
	Image string `json:"Image"`
	Position string `json:"Position"`
	Name string `json:"Name"`	
}

var (
	templatesDir = "templates/"
	t = template.Must(template.ParseGlob(templatesDir + "members.html"))
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")

}

func members(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/members" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
		}
		fmt.Println(r.FormValue("file"))
		AddMember(r.FormValue("name"),r.FormValue("id"), r.FormValue("photo"), r.FormValue("status"))
	}
	a := GetAllowed()
	
	if err := t.ExecuteTemplate(w, "allowed", a); err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(500), 500)
	}
}


func serveW(w http.ResponseWriter, r *http.Request, msg chan *Message) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Client servW: ", err)
		return
	}

	for {
		select {
		case message := <-msg:
			if err = conn.WriteJSON(message); err != nil {
				log.Println("Client serveW <- msg: ", err)
				return
			}

		}
	}
}


func upload(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "", http.StatusBadRequest)
        return
    }

    if err := r.ParseMultipartForm(1024); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    uploadedFile, handler, err := r.FormFile("file")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := handler.Filename

	fileLocation := filepath.Join(dir, "public/static/images", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("done"))

}
