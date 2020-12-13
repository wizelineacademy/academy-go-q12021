package repository

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Record struct {
	ID,
	nombre,
	apellido,
	edad,
	peso,
	estatura string
}

func ReadFile(filepath string) []*Record {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	var records []*Record
	for _, ln := range lines {
		records = append(records, NewRecord(strings.Split(ln, ",")))
	}
	return records
}

func NewRecord(fields []string) *Record {
	return &Record{
		ID:       fields[0],
		nombre:   fields[1],
		apellido: fields[2],
		edad:     fields[3],
		peso:     fields[4],
		estatura: fields[5],
	}
}
