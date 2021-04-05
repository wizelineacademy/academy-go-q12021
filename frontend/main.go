package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"main/config"
	"main/controller"
	"main/model"
)

// Abnormal exit constants
const (
	ExitAbnormalErrorLoadingConfiguration = iota
	ExitAbnormalErrorLoadingCSVFile
)

// Renders a single TechStackItem as html page response
func RenderTechStackItem(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		errorMessage := "Url Param 'id' is missing"
		log.Println(errorMessage)
		fmt.Fprintf(w, "%s", errorMessage)
		return
	}
	// Casting the string number to an integer
	id := keys[0]
	var tmpl *template.Template
	item, err := controller.GetTechStackItems(id)
	if err != nil {
		RenderErrorPage(w)
		return
	}
	tmpl = template.Must(template.ParseFiles("html/item.html"))

	if err := tmpl.Execute(w, item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderErrorPage(w http.ResponseWriter) {
	errorTemplate := template.Must(template.ParseFiles("html/server_error.html"))
	data := model.ErrorPage{
		ErrorTitle: "Internal Server Error",
		Message:    "The server is not responding, please check it's running on the right port or contact support.",
	}
	errorTemplate.Execute(w, data)
}

// Renders a list of TechStackItem as html page response
func RenderTechStackItems(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	items, err := controller.GetItems()
	if err != nil {
		RenderErrorPage(w)
		return
	}
	tmpl = template.Must(template.ParseFiles("html/tech_stack_list.html"))
	data := model.PageData{
		PageTitle:     "My Tech Stack",
		TechStackItem: items,
	}
	tmpl.Execute(w, data)
	controller.WriteDataToCSVFile("result.csv", items)
}

func GetQueryParams(r *http.Request) (queryParams model.QueryParameters) {
	keys := r.URL.Query()

	if val, ok := keys["item_per_workers"]; ok {
		IntItemPerWorkers, err := strconv.Atoi(val[0]) // parse string to int
		if err != nil {
			queryParams.ItemPerWorkers = 1
		} else {
			log.Println("item_per_workers query provided")
			queryParams.ItemPerWorkers = IntItemPerWorkers
		}
	} else {
		log.Println("item_per_workers not provided as query param")
	}

	if val, ok := keys["type"]; ok {
		queryParams.Type = val[0]
	}
	if val, ok := keys["items"]; ok {
		IntItems, err := strconv.Atoi(val[0]) // parse string to int
		if err != nil {
			log.Fatal(err.Error())
			queryParams.Items = 1
		} else {
			queryParams.Items = IntItems
			log.Println("items query provided: value ", IntItems)
		}
	} else {
		queryParams.Items = 1
	}
	return
}

func RenderMovies(w http.ResponseWriter, r *http.Request) {
	// Casting the string number to an integer
	queryParams := GetQueryParams(r)

	response := controller.GetMovies(model.QueryParameters{Items: queryParams.Items, ItemPerWorkers: 1, Type: queryParams.Type})

	tmpl := template.Must(template.ParseFiles("html/movies.html"))
	data := model.Page_AllMovies{
		PageTitle: "Cine+",
		Movies:    response.Data,
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderMovieById(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		errorMessage := "Url Param 'id' is missing"
		log.Println(errorMessage)
		fmt.Fprintf(w, "%s", errorMessage)
		return
	}
	// Casting the string number to an integer
	id := keys[0]
	response := controller.GetMoviesById(id)

	tmpl := template.Must(template.ParseFiles("html/item.html"))

	movieMock := model.Movie{
		ImdbTitleId:   "123123",
		Title:         "Hola  k ase",
		OriginalTitle: "Original title",
		Year:          "109123",
	}

	data := model.Page_MovieDetails{
		PageTitle: "Cine+",
		Movie:     movieMock,
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, response.Data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	var configFile string
	flag.StringVar(
		&configFile,
		"public-config-file",
		"config.yml",
		"Path to public config file",
	)
	flag.Parse()

	// Read config file
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Fatal("Failed to load config: %w", err)
		os.Exit(ExitAbnormalErrorLoadingConfiguration)
	}

	// Second deliverable
	http.HandleFunc("/getTechStack", RenderTechStackItems)
	http.HandleFunc("/getTechStackById", RenderTechStackItem)

	// Third deliverable
	http.HandleFunc("/getMovies", RenderMovies)
	http.HandleFunc("/getMovieById", RenderMovieById)
	fmt.Printf("Web app running succesfully on port [%s].", cfg.HTTPPort)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, nil))
}
