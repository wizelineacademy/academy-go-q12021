package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

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
func TestGetData(t *testing.T) {
	tests := []struct {
		name      string
		wantItems []Item
	}{
		{
			name: "Name", 
			wantItems: []Item{
				{
					Id: "0",
					Title: "javascript",
				},
				{
					Id: "1",
					Title: "python",
				},
				{
					Id: "2",
					Title: "c++",
				},
				{
					Id: "3",
					Title: "go",
				},
				{
					Id: "4",
					Title: "swift",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotItems := GetData(); !reflect.DeepEqual(gotItems, tt.wantItems) {
				t.Errorf("GetData() = %v, want %v", gotItems, tt.wantItems)
			}
		})
	}
}
