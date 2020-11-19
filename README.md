# First Gen Pokedex

Find out your First Gen Pokemon details, search it by the name

## Usage

Refer to `db/README.md` to find out how to initialize database

```sh
Pokedex

Usage:
  go run . csv <csvfile> <PokemonName>
  go run . sqlite3 <dbfile> <PokemonName>
  go run . postgres <PokemonName>
  go run . -h | --help

Options:
  -h --help     Show this screen.
```

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

These are the main requirements we will evaluate:

- Use all that you've learned in the course:
  - Best practices
  - Go basics
  - HTTP handlers
  - Error handling
  - Structs and interfaces
  - Clean architecture
  - Unit testing
  - CSV file fetching

## Getting Started

To get started, follow these steps:

1. Fork this project
1. Make your project private
1. Grant your mentor access to the project
1. Commit periodically
1. Apply changes according to the mentor's comments
1. Have fun!

## Deliverables

We provide the delivery dates so you can plan accordingly; please take this challenge seriously and try to make progress constantly.

It’s worth mentioning that you’ll ONLY get feedback from the review team for your first deliverable, so you will have a chance to fix or improve the code based on our suggestions.

For the final deliverable, we will provide some feedback, but there is no extra review date. If you are struggling with something, contact your mentor and peers to get help on time. Feel free to use the slack channel available.

## First Deliverable (due November 22th 23:59PM)

Based on the self-study material and mentorship covered until this deliverable, we suggest you perform the following:

- Select architecture
- Read a CSV or DB
- Handle Errors for CSV/DB (not valid, missing, incorrect connection)
- Use best practices

> Note: what’s listed in this deliverable is just for guidance and to help you distribute your workload; you can deliver more or fewer items if necessary. However, if you deliver fewer items at this point, you have to cover the remaining tasks in the next deliverable.

## Final Deliverable (due December 13th 23:59PM)

- Store read CSV/DB in a structure
- Loop structure and print data
- Send to own API data read
- Add unit testing
- Refactor

> Important: this is the final deliverable, so all the requirements must be included. We will give you feedback and you will have 3 days more to apply changes. On the third day, we will stop receiving changes at 11:00 am.

## Submitting the deliverables

For submitting your work, you should follow these steps:

1. Create a pull request with your code, targeting the master branch of the repository golang-bootcamp-2020.
2. Fill this [form](https://forms.gle/ogQtHBk6DtZ5yKUM9) including the PR’s url
3. Stay tune for feedback
4. Do the changes according to your mentor's comments

## Documentation

### Must to learn

- [Go Tour](https://tour.golang.org/welcome/1)
- [Go basics](https://www.youtube.com/watch?v=C8LgvuEBraI)
- [Git](https://www.youtube.com/watch?v=USjZcfj8yxE)
- [Tool to practice Git online](https://learngitbranching.js.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [How to write code](https://golang.org/doc/code.html)
- [Go by example](https://gobyexample.com/)
- [Go cheatsheet](http://cht.sh/go/:learn)
- [Any talk by Rob Pike](https://www.youtube.com/results?search_query=rob+pike)
- [The Go Playground](https://play.golang.org/)

### Self-Study Material

- [Golang Docs](https://golang.org/doc/)
- [Constants](https://www.youtube.com/watch?v=lHJ33KvdyN4)
- [Variables](https://www.youtube.com/watch?v=sZoRSbokUE8)
- [Types](https://www.youtube.com/watch?v=pM0-CMysa_M)
- [Functions](https://www.youtube.com/watch?v=feU9DQNoKGE)
- [Error Handling](https://www.youtube.com/watch?v=26ahsUf4sF8)
- [Modules](https://www.youtube.com/watch?v=Z1VhG7cf83M)
  - [Part 1 and 2](https://blog.golang.org/using-go-modules)
- [Go tools](https://dominik.honnef.co/posts/2014/12/an_incomplete_list_of_go_tools/)
- [More Go tools](https://dev.to/plutov/go-tools-are-awesome-bom)
- [Clean Architecture](https://medium.com/@manakuro/clean-architecture-with-go-bce409427d31)
- [For Loops](https://www.youtube.com/watch?v=0A5fReZUdRk)
- [Arrays and Slices](https://www.youtube.com/watch?v=d_J9jeIUWmI)
- [Conditional statements: If](https://www.youtube.com/watch?v=QgBYnz6I7p4)
- [Multiple options conditional: Switch](https://www.youtube.com/watch?v=hx9iHend6jM)
- [Maps](https://www.youtube.com/watch?v=p4LS3UdgJA4)
- [Structures](https://www.youtube.com/watch?v=w7LzQyvriog)
- [Structs and Functions](https://www.youtube.com/watch?v=RUQADmZdG74)
- [Pointers](https://tour.golang.org/moretypes/1)
- [Interfaces](https://tour.golang.org/methods/9)
- [Interfaces](https://gobyexample.com/interfaces)
- [Methods](https://www.youtube.com/watch?v=nYWa5ECYsTQ)
- [Failed requests handling](http://www.metabates.com/2015/10/15/handling-http-request-errors-in-go/)
- [Packages](https://www.youtube.com/watch?v=sf7f4QGkwfE)
- [Unit testing](https://golang.org/pkg/testing/)
- [Functions as values](https://tour.golang.org/moretypes/24)
