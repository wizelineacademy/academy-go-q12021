# Golang Bootcamp

## Introduction

This API can read, add, edit or remove items from a Database. The items are leagues from soccer tournaments with several attributes.

# Requirements

* Go 1.15 or higher

## Installation


1. Clone the project
1. Run the main file: `go run main.go`. This will automatically install all the dependencies, which are gorilla mux and a few for firestore handling.
**Note:** By default port 8081 will be used, but this can be changed by adding the port flag: `go run main.go --port :PORT`




## Setting up the environment

| env variable  | value |
| ------------- |:-------------:|
| api_base      | http://localhost:8081     |

## Get Leagues
A **GET** request to the endpoint **api_base/leagues** will return you all the leagues in the database.

#### Example:

###### Request:
```
curl --location --request GET '{{api_base}}/leagues'
```

###### Response:

```
[
    {
        "country": "Italy",
        "current_season_id": 32523,
        "id": "muQQl8UUP4Y0VFf2x6CY",
        "name": "Serie A"
    },
    {
        "country": "Spain",
        "current_season_id": 32501,
        "id": "yEOmgnPATuWkX0nRsNGD",
        "name": "La Liga",
        "sofascore_id": 8
    },
    {
        "country": "England",
        "current_season_id": 29415,
        "id": "zKuBTQBFa7Dl65FUIlXa",
        "name": "Premier League"
    }
]
```



## Add a League
A **POST** request to the endpoint **api_base/leagues** will add a league with the details you specify in the body of the request.

###### Request Body Requirements:

| Name  | Type |
| ------------- |:-------------:|
| name     | right foo     |
| current_season_id      | int     |


#### Example:

###### Request:
```
curl --location --request POST '{{api_base}}/leagues' \
--data-raw '{
    "country": "Italy",
    "name": "Serie A",
    "current_season_id": 32523
}'
```

###### Response:

```
{
    "country": "Italy",
    "current_season_id": 32523,
    "id": "kyxydckI46Maun5M1RFp",
    "name": "Serie A"
}
```

## Get a League by ID
A **GET** request to the endpoint **api_base/leagues/{id}** will return the requested league details.


#### Example:

###### Request:
```
curl --location --request GET '{{api_base}}/leagues/{{league_id}}'
```

###### Response:

```
{
  "country": "Italy",
  "current_season_id": 32523,
  "id": "muQQl8UUP4Y0VFf2x6CY",
  "name": "Serie A"
}
```



## Update League
A **PATCH** request to the endpoint **api_base/leagues/{id}** will update the requested league details.


#### Example:

###### Request:
```
curl --location --request PATCH '{{api_base}}/leagues/{{league_id}}' \
--data-raw '{
    "current_season_id": 32522
}'
```

###### Response:

```json
{
  "country": "Italy",
  "current_season_id": 32522,
  "id": "muQQl8UUP4Y0VFf2x6CY",
  "name": "Serie A"
}
```



## Delete League
A **DELETE** request to the endpoint **api_base/leagues/{id}** will remove the specified league from the database.


#### Example:

###### Request:
```
curl --location --request DELETE '{{api_base}}/leagues/{{league_id}}'
```

###### Response:

```
{
  "Message": "Succesfully deleted league"
}
```