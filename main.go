// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
)

// set the address to port 8080
var addr = flag.String("addr", ":40000", "http service address")

// serves the home
func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	// if it's not the / path
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	// if its not a get request
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// serve the home file
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()

	// hub.go
	hub := newHub()
	go hub.run()
	// Call our serveHome function when we get a main path (kinda makes some stuff above irrelevant I guess?)
	http.HandleFunc("/", serveHome)
	// When we get a request on the ws path, serveWs
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving websocket!")
		serveWs(hub, w, r)
	})
	// create the listener
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
