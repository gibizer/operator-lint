.PHONY: build
build: test
	go build

.PHONY: clean
clean:
	go clean

.PHONY: test
test:
	go test -v ./...

.PHONY: run
run: build
	go vet -vettool=./operator-lint
