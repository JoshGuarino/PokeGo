test:
	go test ./...

test-verbose:
	go test -v ./...

test-clean:
	go clean -testcache

test-coverage:
	go test -cover ./...

tidy:
	go mod tidy

install:
	go get .
