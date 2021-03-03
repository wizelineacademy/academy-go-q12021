#!/usr/bin/env bash
go build -o pokedex main.go
codesign -s - pokedex