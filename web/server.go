package web

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func StartServer() {
	http.HandleFunc("/", greet)
	log.Println("[!] Server inicializado")
	http.ListenAndServe(":8080", nil)
}
