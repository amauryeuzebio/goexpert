package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("Request iniciado")

	defer log.Println("Request finalizado")

	select {
	case <-time.After(5 * time.Second):
		// Imprime no command line stdout (Server)
		log.Println("Request processado com sucesso")

		// Imprime no browser
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		// Imprime no command line stdout (Server)
		log.Println("Request processado com sucesso")

		// Imprime no browser
		http.Error(w, "Request cancelada pelo cliente", http.StatusRequestTimeout)
	}

}
