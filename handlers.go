package main

import (
	"fmt"
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

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	a := GetAllowed()
//	mem := NewMember("111", "222", "333", "444")
	//m := []*Member{}
	//m = append(m, mem)
	//a := Allowed{Members: m}
	
	fmt.Println("In members", a)
	if err := t.ExecuteTemplate(w, "allowed", a); err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(500), 500)
	}
	//http.ServeFile(w, r, "members.html")
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

