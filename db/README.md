# Data Access Layer

This project uses SQLC to generate type-safe Go code from SQL queries. We have automated the generation of the `sqlc.yaml` configuration file using a shell script. This guide will walk you through the process of adding new SQL query files and schemas, and how to use the script to update the configuration.

## Generating SQLC Configuration

We use a shell script to automate the creation of the `sqlc.yaml` configuration file.

```bash
chmod +x sqlc_config_gen.sh # Make the Script Executable
./sqlc_config_gen # Run the Script
```

## Using the Generated Code

After generating the configuration, run SQLC to generate Go code:

```bash
sqlc generate
```

This will create Go packages in the dal directory for each query group. You can use generated code as following

```go
import (
    "github.com/yourrepo/yourproject/authors"
)

func main() {
  // Example using authors package
	authorQueries := authors.New(conn)
	
  // list all authors
	authorsList, err := authorQueries.List(ctx)

  // create an author
  insertedAuthor, err := authorQueries.Create(ctx, authors.CreateParams{
    Name: "Brian Kernighan",
    Bio:  pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
  })
}
```

## Directory Structure

```bash
├── queries
│   └── authors
│       └── create.query.sql
│       └── list.query.sql
│       └── query.sql
├── sqlc.yaml
└── sqlc_config_gen.sh
```

### Adding a New Query

1. **Create a New Query File**:
   - Navigate to the `queries` directory.
   - Create a new subdirectory if needed (e.g., `publishers`).
   - **NOTE:**
     - Use plural & camelCase for dir name, because this name will be a package name under `dal` package
   - Add a new SQL file for your queries (e.g., `list.query.sql`).
   - **NOTE:** 
     - It's important to end query file with `query.sql`. 
     - You can avoid adding separate file for every query and write all the queries in one file `query.sql`

2. **Write SQL Queries**:
   - Use SQLC annotations to define query names and parameters.
   - Instead of using name `GetPublisherByID` you should use `GetByID`. This is because config file will generate package `publishers` and add all `Publisher` related queries under it. So that you can call your query like `dal.publishers.GetByID` instead of `dal.publishers.GetPublisherByID`

   Example:
   ```sql
   -- name: GetByID :one
   SELECT * FROM publishers WHERE id = $1;

   -- name: List :many
   SELECT * FROM publishers;
   ```

### Migrations

sqlc does not perform database migrations for you. However, sqlc is able to differentiate between up and down migrations. **`sqlc` ignores down migrations when parsing SQL files.** Read more [here](https://docs.sqlc.dev/en/stable/howto/ddl.html)

We will be using `golang-migrate` tool for dealing with migrations. [PostgreSQL tutorial for beginners](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)

Use following command to install tool

```bash 
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

#### Generate migration files

Following command will generate two SQL migration files, one for the `"up"` migration and one for the `"down"` migration, using the specified `title`. 

```sh
go run cmd/main.go migrate create YOUR_MIGRATION_TITLE
```

=== OR ===

```sh
migrate create -ext sql -dir db/migrations -seq YOUR_MIGRATION_TITLE
```

#### Run migrations

```sh
go run cmd/main.go migrate up
```

=== OR ===

When using Migrate CLI we need to pass to database URL. Let's export it to a variable for convenience 

```sh
export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/example?sslmode=disable'
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```
