package main

import (
	"AUThConnect/internal/server"
	"fmt"
)

func main() {

	server := server.NewServer()

	fmt.Println("Server listening on port 8000")
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
