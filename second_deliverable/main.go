package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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

func GetItems() (items []Item) {
	// Get the http reponse from api localhost:8080 (first_deliverable)
	resp, err := http.Get("http://localhost:8080/getLanguages")
	if err != nil {
		log.Fatalf(err.Error())
		items = []Item{{Title: "", Id:""}}
		defer resp.Body.Close()
		return
	} else {
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
}

func GetItemsById(id string) (item Item) {
	// Get the http reponse from api localhost:8080 (first_deliverable)
	var url string = "http://localhost:8080/getLanguageById?id=" + id
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("\n\tResponse status:", resp.Status)

	// Print the first 5 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
    	json.Unmarshal([]byte(scanner.Text()), &item) // items slice
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return item
}


func RenderItem(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
    if !ok || len(keys[0]) < 1 {
		errorMessage := "Url Param 'id' is missing"
		log.Println(errorMessage)
		fmt.Fprintf(w, "%s", errorMessage)
        return
    }
	// Casting the string number to an integer
    id := keys[0]
	item := GetItemsById(id)

	tmpl := template.Must(template.ParseFiles("html/item.html"))

	if err := tmpl.Execute(w, item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}	
}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	items := GetItems()
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	data := PageData{
		PageTitle: "My Tech Stack",
		Items: items,
	}
	tmpl.Execute(w, data)
	WriteDataToCSVFile("result.csv", items)
}

func WriteDataToCSVFile(fileName string, items []Item){
	log.Println("Data: ", items)

	csvfile, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("Error creating file csv", err)
	}
	var writter *csv.Writer = csv.NewWriter(csvfile)

    for _, item := range items {
		strSlice := []string{item.Id,item.Title}
    	fmt.Println(strSlice)
		writter.Write(strSlice)
	}		
	// Write any buffered items data to the underlying writer (standard output).
	writter.Flush()

	if err := writter.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

}

func main() {
	log.Printf("\n\tRunning webapp on port %v!", port)
    http.HandleFunc("/", GetAllItems)
	http.HandleFunc("/getItem",RenderItem)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}



