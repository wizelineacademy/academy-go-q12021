package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)
var port string = "3000"
type ItemJson struct {
	Id string `json:"id"`
    Title string `json:"title"`
}

type Item struct {
	Id string 
    Title string
}


type PageData struct {
    PageTitle string
    Items     []Item
}

func GetData() (items []Item) {
	// Get the http reponse from api localhost:8080 (first_deliverable)
	resp, err := http.Get("http://localhost:8080/getLanguages")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("\n\tResponse status:", resp.Status)

	// Print the first 5 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
    	json.Unmarshal([]byte(scanner.Text()), &items) // items slice
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return items
}

func RenderItem(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "%v", "List of Programming languages")
	items := GetData()
	tmpl := template.Must(template.ParseFiles("html/item.html"))
    // for _, item := range items {
	if err := tmpl.Execute(w, items[0]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}	
	// }
}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	data := PageData{
		PageTitle: "My Tech Stack",
		Items: GetData(),
	}
	tmpl.Execute(w, data)
}


func main() {
	log.Printf("\n\tRunning webapp on port %v!", port)
    http.HandleFunc("/", GetAllItems)
	http.HandleFunc("/getItem",RenderItem)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}



