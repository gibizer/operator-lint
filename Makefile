.PHONY: build
build: test
	go build

.PHONY: clean
clean:
	rm plugin.so

.PHONY: test
test:
	go test ./...

.PHONY: run
run: build
	go vet -vettool=./envtest-linter
