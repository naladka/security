package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/huin/goserial"
	"github.com/ivan-bogach/fileserver"
)

var addr = flag.String("addr", ":8080", "http service address")

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
}

func write(msg chan *Message, str string) {
	msg <- &Message{Message: str}
}



func main() {
	GetAllowed()
	msg := make(chan *Message)

	hub := newHub()
	go hub.run(msg)
	var prevAMsg string
	go func() {

		buf := make([]byte, 16)
		d := &Device{hub: hub, send: make(chan []byte), localAddr: "loc"}
		d.hub.reg <- d

		gays := GetAllowed()

		c := &goserial.Config{Name: findArduino(), Baud: 115200}
		s, err := goserial.OpenPort(c)
		if err != nil {
			log.Fatal(err)
		}

		for {
			n, err := io.ReadFull(s, buf)
			if err != nil {
				log.Fatal(err)
			}

			aMsg := string(buf[:n])
			fmt.Println(prevAMsg + "<===>" + aMsg)
			fmt.Println(prevAMsg == aMsg)
			if prevAMsg == aMsg {
				fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAA")
				continue
			}
			prevAMsg = aMsg
			btns := strings.Split(aMsg, "")[0]
			lastCardID := strings.Split(aMsg, "")[1:]
			fmt.Printf("%s\n", aMsg)
			switch btns {
			case "1":
				go write(msg, "ОТКРЫТА ДВЕРЬ!")
				_, err := s.Write([]byte("0"))
				if err != nil {
					log.Fatal(err)
				}
			case "2":
				go write(msg, "ЗВОНОК")
			case "3":
				go write(msg, "ЗВОНОК И ДВЕРЬ ОТКРЫТА")
			case "4":
				go write(msg, "ПОСТОРОННИЕ!")
				
			case "5":
				go write(msg, "КОД И ДВЕРЬ")
			case "6":
				go write(msg, "ЗВОНОК И КОД")
			case "7":
				fmt.Println("Together")
			}

			idFromArd := strings.Join(lastCardID, "")
			//fmt.Println("Last time come in: %s", gays[strings.Join(lastCardID, "")])
			//if _, ok := gays[strings.Join(lastCardID, "")]; ok {
			if gays.IDExist(idFromArd) {
				n, err := s.Write([]byte("1"))
				fmt.Println("wroten %d bytes", n)
				go write(msg, "Пришел " +  gays.GetName(idFromArd))
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	flag.Parse()
	
	http.Handle("/static/", http.StripPrefix("/static/", fileserver.FileServer(
		http.Dir("./public/static"),
		fileserver.FileServerOptions{
				IndexHTML: true,
		},
	)))
	http.HandleFunc("/members", members)
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveW(w, r, msg)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


func findArduino() string {
	contents, _ := ioutil.ReadDir("/dev")

	// Look for what is mostly likely the Arduino device
	for _, f := range contents {
		if strings.Contains(f.Name(), "tty.usbserial") ||
			strings.Contains(f.Name(), "ttyACM0") {
			return "/dev/" + f.Name()
		}
	}

	// Have not been able to find a USB device that 'looks'
	// like an Arduino.
	return ""
}