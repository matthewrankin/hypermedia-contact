package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func contactsFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Contacts\n")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/contacts", contactsFunc)
	fmt.Println("Starting the server on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
