package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

type Contact struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	Phone     string `json:"phone"`
}

func getContacts() []Contact {
	resp, err := http.Get("https://61116372c38a0900171f11c0.mockapi.io/contacts")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var cs []Contact
	err = json.Unmarshal(body, &cs)
	if err != nil {
		fmt.Println("error:", err)
	}

	return cs
}

func main() {
	c := getContacts()

	h := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello world!\n")
	}

	handlerRandom := func(w http.ResponseWriter, _ *http.Request) {
		randomIndex := rand.Intn(len(c) - 1)
		randomContact := c[randomIndex]

		json.NewEncoder(w).Encode(randomContact)
	}

	http.HandleFunc("/", h)
	http.HandleFunc("/contacts", handlerRandom)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
