# convoy    ![Go](https://github.com/javiertlopez/golang-bootcamp-2020/workflows/Go/badge.svg?branch=main)  [![codecov](https://codecov.io/gh/javiertlopez/golang-bootcamp-2020/branch/main/graph/badge.svg?token=8XM6UUO8UZ)](https://codecov.io/gh/javiertlopez/golang-bootcamp-2020)

    convoy espacial, que tan lejos nos llevará?

## Overview 

The app handles events such as weddings, family trips and school excursions; it also manages hotel reservations for its guests.

## What's the deal?

The app is based on Clean Architecture. 

- Domain (under the `model` folder)
- Usecase (under the `usecase` folder)
- Controller (under the `controller` folder)
- Database/Cache (under the `repository` folder)
- Interface (under the `router` folder)

Events and Reservations are stored in Mongo Atlas. Reservations are cached in a CSV file.

## Development

### Install dependencies

Install dependencies using
```bash
go get ./...
```

## Run app

### Docker

Create a docker image

```bash
docker build -t javiertlopez/convoy .
```

Run the docker container

```bash
docker run --env-file=dev.env -d -p 8080:8080 javiertlopez/convoy
```

**Note.** It is required to fill the `dev.env` file.

### From terminal
Export environment environment variable with **Mongo Atlas Connection String**:

```bash
export ADDR="0.0.0.0:8080"
export CSV_FILE="cache.csv"
export MONGO_STRING="connectionString"
export MONGO_DB="databaseName"
```

Run app directly from terminal:

```bash
go run .
```

### From Visual Studio

The project is setup to run directly from Visual Studio Run tab. Just fill the credentials, and rename the file from `dev.env.example` to `dev.env`.

## Endpoints

### Get Event

`GET /events/{id}`

Example 1. Event without reservations:

```bash
curl --location --request GET 'http://127.0.0.1:8080/events/a9f99619-13e4-4488-9c7f-af9e311538a1'
```

Response:
```json
{
    "id": "a9f99619-13e4-4488-9c7f-af9e311538a1",
    "description": "Christmas Eve",
    "type": "private",
    "status": "created",
    "created_at": "2020-11-27T21:14:14.669Z",
    "updated_at": "2020-11-27T21:14:14.669Z",
    "event_date": "2006-01-25T01:00:00Z",
    "event_location": "Rosewood San Miguel de Allende",
    "name": "Hiram Torres",
    "phone": "3333333333",
    "email": "hiram@jtl.mx"
}
```

Example 2. Event with reservations:

```bash
curl --location --request GET 'http://127.0.0.1:8080/events/ddba97c2-bfd5-46df-85f3-b72e75e9efd0'
```

Response:
```json
{
    "id": "ddba97c2-bfd5-46df-85f3-b72e75e9efd0",
    "description": "New Year's Eve",
    "type": "private",
    "status": "created",
    "created_at": "2020-11-27T21:19:47.022Z",
    "updated_at": "2020-11-27T21:19:47.022Z",
    "event_date": "2020-12-31T01:00:00Z",
    "event_location": "Rosewood San Miguel de Allende",
    "name": "Hiram Torres",
    "phone": "33333333",
    "email": "hiram@jtl.mx",
    "reservations": [
        {
            "id": "97b658e2-2f3c-43f9-94ea-915c633660cf",
            "status": "created",
            "plan": "european",
            "adults": 2,
            "minors": 0,
            "adult_fee": 4999,
            "minor_fee": 0,
            "arrival": "2020-12-31T01:00:00Z",
            "departure": "2020-12-31T01:00:00Z",
            "name": "Hiram Torres",
            "phone": "33333333",
            "email": "hiram@jtl.mx"
        },
        {
            "id": "729f1bb1-189a-405c-b862-39ec521ba29e",
            "status": "created",
            "plan": "european",
            "adults": 1,
            "minors": 0,
            "adult_fee": 4999,
            "minor_fee": 0,
            "arrival": "2020-12-31T01:00:00Z",
            "departure": "2020-12-31T01:00:00Z",
            "name": "Briseida Romero",
            "phone": "33333333",
            "email": "hiram@jtl.mx"
        }
    ]
}
```

### Create event

`POST /events`

Example 1. Create event without reservations

```bash
curl --location --request POST 'http://127.0.0.1:8080/events' \
--header 'Content-Type: application/json' \
--data-raw '{
    "description": "Christmas Eve",
    "type": "private",
    "status": "created",
    "event_date": "2020-01-24T18:00:00-07:00",
    "event_location": "Rosewood San Miguel de Allende",
    "name": "Hiram Torres",
    "email": "hiram@jtl.mx",
    "phone": "3333333333"
}'
```

Response:

```json
{
    "id": "a9f99619-13e4-4488-9c7f-af9e311538a1",
    "description": "Christmas Eve",
    "type": "private",
    "status": "created",
    "created_at": "2020-11-27T21:14:14.669Z",
    "updated_at": "2020-11-27T21:14:14.669Z",
    "event_date": "2006-01-25T01:00:00Z",
    "event_location": "Rosewood San Miguel de Allende",
    "name": "Hiram Torres",
    "phone": "3333333333",
    "email": "hiram@jtl.mx"
}
```

Example 2. Create event with reservations

```bash
curl --location --request POST 'http://127.0.0.1:8080/events' \
--header 'Content-Type: application/json' \
--data-raw '{
    "description": "New Year'\''s Eve",
    "type": "private",
    "status": "created",
    "event_date": "2020-12-30T18:00:00-07:00",
    "event_location": "Rosewood San Miguel de Allende",
    "reservations": [
        {
            "status": "created",
            "plan": "european",
            "adults": 2,
            "adult_fee": 4999,
            "minors": 0,
            "minor_fee": 0,
            "arrival": "2020-12-30T18:00:00-07:00",
            "departure": "2021-01-02T18:00:00-07:00",
            "name": "Hiram Torres",
            "email": "hiram@jtl.mx",
            "phone": "3333333333"
        },
        {
            "status": "created",
            "plan": "european",
            "adults": 1,
            "adult_fee": 4999,
            "minors": 0,
            "minor_fee": 0,
            "arrival": "2020-12-30T18:00:00-07:00",
            "departure": "2021-01-02T18:00:00-07:00",
            "name": "Briseida Romero",
            "email": "hiram@jtl.mx",
            "phone": "3333333333"
        }
    ],
    "name": "Hiram Torres",
    "email": "hiram@jtl.mx",
    "phone": "3333333333"
}'
```

Response:
```json
{
    "id": "ddba97c2-bfd5-46df-85f3-b72e75e9efd0",
    "description": "New Year's Eve",
    "type": "private",
    "status": "created",
    "created_at": "2020-11-27T21:19:47.022Z",
    "updated_at": "2020-11-27T21:19:47.022Z",
    "event_date": "2020-12-31T01:00:00Z",
    "event_location": "Rosewood San Miguel de Allende",
    "name": "Hiram Torres",
    "phone": "33333333",
    "email": "hiram@jtl.mx",
    "reservations": [
        {
            "id": "97b658e2-2f3c-43f9-94ea-915c633660cf",
            "status": "created",
            "plan": "european",
            "adults": 2,
            "minors": 0,
            "adult_fee": 4999,
            "minor_fee": 0,
            "arrival": "2020-12-31T01:00:00Z",
            "departure": "2020-12-31T01:00:00Z",
            "name": "Hiram Torres",
            "phone": "33333333",
            "email": "hiram@jtl.mx"
        },
        {
            "id": "729f1bb1-189a-405c-b862-39ec521ba29e",
            "status": "created",
            "plan": "european",
            "adults": 1,
            "minors": 0,
            "adult_fee": 4999,
            "minor_fee": 0,
            "arrival": "2020-12-31T01:00:00Z",
            "departure": "2020-12-31T01:00:00Z",
            "name": "Briseida Romero",
            "phone": "33333333",
            "email": "hiram@jtl.mx"
        }
    ]
}
```

## Testing

Run tests
```bash
go test -v ./...
```
