# Golang Bootcamp Project

## Use Cases

### Fetch a quote from an external source.

**Name**: Quote fetch endpoint.

**Description**: Actor fetches a new quote, it becomes stored automatically.

**Actors**: User.

**Input**: Empty endpoint call.

**Output**: Quote in JSON format, as it becomes stored in the database as well.

**Basic Flow**: 

1. User makes an API call through some standard http client.
2. Fetch a new quote from quote garden.
3. Store the quote in the database (CSV file).
4. Return the quote to the user.

### Read all the stored quotes.

**Name**: Get all stored quotes endpoint.

**Description**: Actor fetches all stored quotes.

**Actors**: User.

**Input**: Empty endpoint call.

**Output**: Quotes as an array inside a JSON object.

**Basic Flow**:

1. User makes an API call through some standard http client.
2. Fetch all stored quotes from database.
3. Return all the fetched quotes to the user.

# Implementation details

## API Endpoints Declaration

| Name                 | Description                        |
| -------------------- | ---------------------------------- |
| [**POST**] /v0/quote | Stores a new quote and returns it. |
| [**GET**] /v0/quote  | Returns all quotes from database.  |

## Architecture

Clean architecture, the items are declared this way:

| Package Name        | Equivalent to            | Role                                                         |
| ------------------- | ------------------------ | ------------------------------------------------------------ |
| entity              | Entities Layer           | Business object model layer. Contains structures and also the use cases interfaces. |
| integration         | External interfaces      | Database and web access layer.                               |
| integration/router  | Web interface            | HTTP API.                                                    |
| service             | Interface Adapters Layer | Use case input/output port implementations layer.            |
| service/config      | Storage                  | Loads configuration.                                         |
| service/handler     | Controller               | Recieves an API request.                                     |
| service/quotegarden | Gateway                  | Fetches a new quote from quote garden, REST API Client.      |
| service/repository  | Storage                  | Stores and reads from quotes database.                       |
| usecase             | Use Cases Layer          | Business rules declaration layer.                            |
| usecase/interactor  | Quote Use Case           | Coordinates the available services according to the business rules.|

