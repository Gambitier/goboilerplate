#!/bin/bash

set -e

# Base config parts
version="version: \"2\""
schema="    schema: \"./db/migrations\""
sql_package="    sql_package: \"pgx/v5\""

# Start the config file
echo "$version" > sqlc.yaml
echo "sql:" >> sqlc.yaml

# Find all directories under queries/
for dir in ./db/queries/*/
do
    # Remove trailing slash
    dir=${dir%/}
    # Extract the directory name (e.g., author)
    name=$(basename "$dir")

    # Create the config block for this directory
    cat <<EOL >> sqlc.yaml
  - engine: "postgresql"
    queries: "$dir/*.sql"
    schema: "./db/migrations"
    gen:
      go:
        package: "$name"
        out: "./db/dal/$name"
        sql_package: "pgx/v5"
EOL

done
