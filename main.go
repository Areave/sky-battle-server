package main

import (
	"fmt"
	"net/http"
	"os"
)

// простой обработчик
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go server on Fly.io, bitch!\n")
}

func main() {
	// регистрируем обработчик
	http.HandleFunc("/", handler)

	// Fly.io задаёт порт через переменную среды PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // локально
	}

	fmt.Println("Server is listening on port:", port)

	// запускаем сервер
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
