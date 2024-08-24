# gocomm

## Running the app

### Hot Reload

[Air](https://github.com/air-verse/) is yet another live-reloading command line utility for developing Go applications. 

```bash
go install github.com/air-verse/air@latest
```

once air utility is installed, simply run

```bash
air
```

## Guides

- [Data Access Layer](./db/README.md)
- [Image Processing Setup](./imageProcessor/)

## Usage

### Config 

* Copy `sample.config.json` & create `config.json` file in the same directory
* Update config values in `config.json` 

### gRPC

```bash
protoc --go_out=. --go-grpc_out=. **/*.proto
```

Starting grpcui tool

**Note**: For this you need to set env variable `environment` to `development`

```bash
grpcui -plaintext localhost:3001
```

### Wire

Installation

```bash
go install github.com/google/wire/cmd/wire@latest
```

Use following command to create wire generated files across all the directories in the project.

```bash
wire ./...
```

### Swagger

Install swag by using:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

generate swagger docs

```bash
go generate ./...
```

### sqlc

sqlc is distributed as a single binary with zero dependencies.

#### installation

macOS
```bash
brew install sqlc
```

Ubuntu
```bash
sudo snap install sqlc
```

go install (Go >= 1.17)
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

#### usage

generate `sqlc` generated files

```bash
sqlc generate ./...
```
