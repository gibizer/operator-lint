.PHONY: plugin
plugin: test
	go build -buildmode=plugin golangci-lint-plugin/plugin.go

.PHONY: clean
clean:
	rm plugin.so

.PHONY: test
test:
	go test ./...

