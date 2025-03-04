test:
	go test ./...

test verbose:
	go test -v ./...

tidy:
	go mod tidy

install:
	go get .
