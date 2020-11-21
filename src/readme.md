# Rick and morty API
Challenge from wizeline golang bootcamp 2020

This API fetch data from https://rickandmortyapi.com/ following best practices and clean architecture.
The idea of this project is fetch data from rickandmortyapi, clean not necessary information and store in a CSV file as backup.
Read data from this file and process in a go map in order to reduce the complexity of searches to O(1)
in case of search by id. Also, store a CSV that relates id with names to give the user posibility
to obtain the id from the character's name. 

## Get app status
### `/health`
Confirm that app is up and running

## Fetch data from API
### `/data/fetch`
Fetch data from rickandmortyapi. As optional query param, you can pass the max
number of pages to fetch as maxPages e.g. `?maxPages=2`  

## Get all characters from backup (CSV)
### `/api/characters`
Get characters from go map. Initially, the data is read from CSV file (if exists) and store the data in a go map
to process faster (since the file is only read in this moment or in a fetch)

## Get character by id
### `/api/character/:id` 
Get character given an id. Must be a valid not empty string 

## Get character by name
### `/api/findId/:name` 
Get character id given a complete name. Must be a valid not empty string 