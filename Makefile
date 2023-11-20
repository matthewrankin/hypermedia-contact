include config/.env.dev
export

name = "foo"
.PHONY : help check lint cover local dev docker dock
.PHONY : up down dbup dbdown mig

help:
	@echo "You can perform the following:"
	@echo ""
	@echo "  check         Format, vet, and unit test Go code"
	@echo "  cover         Show test coverage in html"
	@echo "  lint          Lint Go code using staticcheck"
	@echo "  local         Build for local operating system"
	@echo "  dev           Build and run for local operating system"
	@echo "  docker        Build in Docker container"
	@echo "  dock          Build and run in Docker container"
	@echo "  up            Start dev API & DB containers using Docker Compose"
	@echo "  down          Stop dev API & DB containers using Docker Compose"
	@echo "  dbup          Run DB migrate up"
	@echo "  dbdown        Run DB migrate down"
	@echo "  dbseed        Run DB migrations to seed DB with data"
	@echo "  mig           Make new migration name=..."

check:
	@echo "Formatting, vetting, and testing Go code"
	go fmt ./...
	go vet ./...
	go test ./... -cover

lint:
	@echo "Linting code using staticcheck"
	staticcheck -f stylish ./...

cover:
	@echo "Test coverage in html"
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

local:
	@echo "Building for local operating system"
	env go build -o dist/snippetbox ./cmd/web/

dev: local
	dist/snippetbox -addr=":4100"

docker:
	@echo "Building in Docker container"
	docker build -t snippetbox:latest .

dock: docker
	@echo "Building and running in Docker container"
	docker run --rm -p 4100:4100 snippetbox:latest

up:
	docker compose build
	docker compose up

down:
	docker compose down

dbup:
	migrate -database ${POSTGRES_URL} -path migrations up

dbdown:
	migrate -database ${POSTGRES_URL} -path dbseed down
	migrate -database ${POSTGRES_URL} -path migrations down

dbseed:
	migrate -database ${POSTGRES_URL} -path dbseed up

mig:
	@echo "Creating new migration"
	migrate create -seq -ext=.sql -dir=./migrations $(name)
