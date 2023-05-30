package main

import (
	"group/handle"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handle.MainPage)
	mux.HandleFunc("/artist", handle.ArtistPage)
	fileserver := http.FileServer(http.Dir("./ui/style/"))
	mux.Handle("/resources/", http.StripPrefix("/resources", fileserver))
	log.Println("Запуск веб-сервера на http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
