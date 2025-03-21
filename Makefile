test:
	go test ./...

test-verbose:
	go test -v ./...

test-clean:
	go clean -testcache

tidy:
	go mod tidy

install:
	go get .
