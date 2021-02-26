package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnmarshallSearchResponse(t *testing.T) {
	var searchData SearchResponse
	testData := "{\"numFound\": 100,\"start\": 0,\"docs\": [{ \"cover_i\": 855045,\"has_fulltext\": false,\"title\": \"Untangling Tolkien\",\"title_suggest\": \"Untangling Tolkien\",           \"type\": \"work\",            \"ebook_count_i\": 0,            \"edition_count\": 1,            \"key\": \"/works/OL5747274W\",\r\n            \"last_modified_i\": 1383146561,\r\n            \"cover_edition_key\": \"OL3319316M\",\r\n            \"first_publish_year\": 2003\r\n}]}"
	json.Unmarshal([]byte(testData), &searchData)
	fmt.Println(searchData.Start)
	fmt.Println(searchData.NumFound)
	fmt.Println(len(searchData.Doc))
	// fmt.Println(searchData.Doc[0].Key)

	if searchData.Start != 0 {
		t.Error("Expected 0 in start")
	}
}
