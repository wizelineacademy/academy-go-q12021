package main

import (
	"testing"
)

var Items []Item = []Item{
	{Id: "1", Title: "Netflix"},
	{Id: "2", Title: "Dispney+"},
	{Id: "3", Title: "HBO Max"},
	{Id: "4", Title: "Paramount+"},
	{Id: "5", Title: "Universal+"},
}

func TestWriteDataToCSVFile(t *testing.T) {
	type args struct {
		fileName string
		items    []Item
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				fileName: "test.csv", 
				items: Items,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteDataToCSVFile(tt.args.fileName, tt.args.items)
		})
	}
}
