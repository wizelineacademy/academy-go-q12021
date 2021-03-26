package main

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetData(t *testing.T) {
	tests := []struct {
		name      string
		wantItems []Item
	}{
		{
			name: "Name",
			wantItems: []Item{
				{
					Id:    "1",
					Title: "1",
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

func TestRenderItem(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RenderItem(tt.args.w, tt.args.r)
		})
	}
}
