package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Quote = struct {
	ID         string
	NationalID string
	Name       string
	Title      string
	Message    string
	Pokedex    string
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	view := `
		<h1>Welcome to the PokéLimbo!</h1>
		<p>
			Feel free to go to the
			<a title="/pokehell" href="/pokehell">PokéHell</a>
			... if you dare.
		</p>
	`
	log.Print("Hit: Home")
	fmt.Fprintf(w, view)
}

func pokehell(w http.ResponseWriter, r *http.Request) {
	log.Print("Hit: PokéHell")
	id := r.URL.Query().Get("id")
	db, err := os.Open("db.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(db)
	reader.Comment = '#'
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var quote Quote

	var quotesList []Quote

	for _, currentQuote := range lines {
		quote.ID = currentQuote[0]
		quote.NationalID = currentQuote[1]
		quote.Name = currentQuote[2]
		quote.Title = currentQuote[3]
		quote.Message = currentQuote[4]
		quote.Pokedex = currentQuote[5]

		quotesList = append(quotesList, quote)
	}

	if id == "" {
		log.Print("ID: None")
		jsonData, _ := json.Marshal(quotesList)
		fmt.Fprintf(w, string(jsonData))
	} else {
		log.Print("ID: " + id)
		found := false
		for _, quote := range quotesList {
			if quote.ID == id {
				found = true
				jsonData, _ := json.Marshal(quote)
				log.Print("Found: " + string(jsonData))
				fmt.Fprintf(w, string(jsonData))
			}
		}
		if found == false {
			log.Print("Found: None")
			jsonData, _ := json.Marshal([]string{})
			fmt.Fprintf(w, string(jsonData))
		}
	}

}

func requestHandlers() {
	http.HandleFunc("/", home)
	http.HandleFunc("/pokehell", pokehell)
}

func main() {
	requestHandlers()
	log.Print("Server running to PokéHell on port 666")
	log.Fatal(http.ListenAndServe(":666", nil))
}
