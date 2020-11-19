# Database Support

## Using SQLite3

Database can be generated following the next commands:
```
cd {project_root}
sqlite3 {db_name} < ./resources/pokemon.sql
```

## Using PostgreSQL

The source file is located at `{project_root}/resources/pokemon.sql` there is a docker compose file which can be used to start a database instance:
```
cd {project_root}/resources
docker-compose up -d
```

## Interesting packages:
* [**database/sql**](https://golang.org/pkg/database/sql/) This is the standard library in go
* [**jmoiron/sqlx**](https://pkg.go.dev/github.com/jmoiron/sqlx) Extension of the Go database/sql library
* [**Masterminds/squirrel**](https://pkg.go.dev/github.com/Masterminds/squirrel) SQL Generator
* [**gocraft/dbr**](https://pkg.go.dev/github.com/gocraft/dbr) Midway between Standard Library and ORM builds on top of sqlx and squirrel
* [**gorm.io/gorm**](https://pkg.go.dev/gorm.io/gorm) Fully fledged ORM database support
