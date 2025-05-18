# Golang Check Vuln Dependencies
code-quality:
	go mod verify
	go tool govulncheck ./...
	go tool golangci-lint run ./...

# Golang Start
start:
	go run cmd/server/main.go

dev:
	go tool air

# Golang Update Dependencies
update:
	go mod tidy

# Golang Clean Cache
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



# Docs
docs-install-depends:
	cd docs && npm i

docs-compile-watch:
	cd docs && npm run watch

docs-preview:
	cd docs && npm run preview
