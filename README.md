# convoy

    convoy espacial, que tan lejos nos llevará?

## Overview 

The app handles events such as weddings, family trips and school excursions; it also manages hotel reservations for its guests.

## What's the deal?

The app is based on the Clean Architecture. The first deliverable contains:

- Domain (under the `model` folder)
- Usecase (under the `usecase` folder)
- Controller (under the `controller` folder)
- Database (under the `repository` folder)
- Interface (under the `router` folder)

## Development

### Install dependencies

Install dependencies using
```bash
go get
```

## Run app

### From terminal
Export environment environment variable with **Mongo Atlas Connection String**:

```bash
export ADDR="0.0.0.0:8080"
export MONGO_STRING="connectionString"
export MONGO_DB="databaseName"
```

Run app directly from terminal:

```bash
go run .
```

### From Visual Studio

The project is setup to run directly from Visual Studio Run tab. Just fill the credentials, and rename the file from `dev.env.example` to `dev.env`.
