package repository

import (
	"bufio"
	"log"
	"os"
)

type Record struct {
	ID,
	nombre,
	apellido,
	edad,
	peso,
	estatura string
}

func ReadFile(file string) []string {
	file, err := os.Open(file)

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	var records []string
	for _, ln = range text {
		records = append(records, NewRecord(strings.split(ln)))
	}
	return records
}

func NewRecord(fields *[]string) *Record {
	return &Record{
		ID:       fields[0],
		nombre:   fields[1],
		apellido: fields[2],
		edad:     fields[3],
		peso:     fields[4],
		estatura: fields[5],
	}
}
