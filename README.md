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

## Testing

Run tests
```bash
go test -v ./...
```
