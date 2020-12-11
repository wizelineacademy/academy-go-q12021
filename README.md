# Golang Bootcamp

## The Challenge

The purpose of the challenge is for you to demonstrate your Golang skills. This is your chance to show off everything you've learned during the course!!

You will build and deliver a whole Golang project on your own. We don't want to limit you by providing some fill-in-the-blanks exercises, but instead request you to build it from scratch.
We hope you find this exercise challenging and engaging.

The goal is to build a REST API which must include:

- An endpoint for reading from an external DB or API
  - Write the information in a CSV file
- An endpoint for reading the CSV
  - Display the information as a JSON
- Unit testing for the principal logic
- Follow conventions, best practices
- Clean architecture

## Requirements

- Go 1.12.9 or avobe
- github.com/gorilla/mux
- github.com/lib/pq
- github.com/stretchr/testify/assert
- github.com/go-sql-driver/mysql
- database/sql
- docker

## Getting Started

Run the next make commands:

`make setup-local` to configure the docker db and install github images

`go run main.go` to run the project after configuration

`go test` run tests
