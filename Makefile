# Makefile for gocomm

# Variables
AIR_CMD=air
PROTOC_CMD=protoc
GRPCUI_CMD=grpcui -plaintext localhost:3001
WIRE_CMD=wire ./...
SWAG_CMD=swag
SQLC_CMD=sqlc generate ./...
CONFIG_FILE=config.json
SAMPLE_CONFIG_FILE=sample.config.json

# ANSI color codes
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[1;33m
BLUE=\033[0;34m
MAGENTA=\033[0;35m
CYAN=\033[0;36m
WHITE=\033[1;37m
NC=\033[0m # No Color

# Targets

.PHONY: all
all: help

.PHONY: help
help:
	@echo "Usage:"
	@echo "  make dev-init            Set up development environment"
	@echo "  make host-ip             Display host ip to be used in config file"
	@echo "  make migrate-db          Applys `up` database migrations"
	@echo "  make run                 Run the application with hot reload"
	@echo "  make build               Build source"
	@echo "  make install-tools       Install all required tools"
	@echo "  make grpc                Generate gRPC code"
	@echo "  make grpcui              Start grpcui tool"
	@echo "  make wire                Generate wire files"
	@echo "  make swagger             Generate Swagger documentation"
	@echo "  make sqlc                Generate SQLC code"
	@echo "  make config              Create and update config file"
	@echo "  make clean               Clean generated files"
	@echo "  make build-log           Generate Makefile build log"

.PHONY: dev-init
dev-init: install-tools config swagger sqlc wire build host-ip
	@echo "${GREEN}Development environment setup complete.${NC}"
	@echo "${YELLOW}Please update config values in $(CONFIG_FILE) and run 'make run' to start the application.${NC}"

.PHONY: build
build: 
	go mod tidy
	go build

.PHONY: migrate-db
migrate-db: 
	go run cmd/main.go migrate up

.PHONY: host-ip
host-ip:
	$(eval HOST_IP := $(shell ip route | grep default | awk '{print $$3}'))
	@echo "=================================="
	@echo "${CYAN}HOST_IP is $(HOST_IP)${NC}"
	@echo "=================================="

.PHONY: install-tools
install-tools:
	go install github.com/air-verse/air@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: run
run:
	$(AIR_CMD)

.PHONY: grpc
grpc:
	$(PROTOC_CMD) --go_out=. --go-grpc_out=. **/*.proto

.PHONY: grpcui
grpcui:
	$(GRPCUI_CMD)

.PHONY: wire
wire:
	$(WIRE_CMD)

.PHONY: swagger
swagger:
	go generate ./...

.PHONY: sqlc
sqlc:
	$(SQLC_CMD)

.PHONY: config
config:
	cp $(SAMPLE_CONFIG_FILE) $(CONFIG_FILE)
	@echo "Please update config values in $(CONFIG_FILE)"

.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	rm -rf _apidocs
	rm -rf tmp

.PHONY: build-log
build-log:
	make --dry-run --always-make --keep-going --print-directory > Makefile-build.log
