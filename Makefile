# Golang Check Vuln Dependencies
code-quality:
	go mod verify
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run ./...
	go run github.com/securego/gosec/v2/cmd/gosec@latest ./...

# Golang Start
start:
	go run cmd/server/main.go

# Golang Update Dependencies
update:
	go mod tidy

# Golang Clean Cache
clean: clean-go-cache clean-test-cache
	rm -rf ./bin

clean-go-cache:
	go clean -cache

clean-test-cache:
	go clean -testcache


# Test Targets
unittests:
	go clean -testcache && go test ./test/unittests/...

unittests-verbose:
	go clean -testcache && go test -v ./test/unittests/...

unittests-coverage:
	go clean -testcache && go test -v -coverpkg=./... -coverprofile=coverage.out ./test/unittests/...
	go tool cover -html=coverage.out


# Docker Compose Targets
compose-up:
	docker compose up -d

compose-down:
	docker compose down

compose-up-db:
	docker compose up -d db 

compose-down-db:
	docker compose down db