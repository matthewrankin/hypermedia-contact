help:
	@echo "You can perform the following:"
	@echo ""
	@echo "  check         Format, vet, and unit test Go code"
	@echo "  cover         Show test coverage in html"
	@echo "  lint          Lint Go code using staticcheck"
	@echo "  local         Build for local operating system"

check:
	@echo 'Formatting, vetting, and testing Go code'
	go fmt ./...
	go vet ./...
	go test ./... -cover

lint:
	@echo 'Linting code using staticcheck'
	staticcheck -f stylish ./...

cover:
	@echo 'Test coverage in html'
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

local:
	@echo 'Building for local operating system'
	env go build -o dist/contacts
