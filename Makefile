src = $(shell find . -type f -name '*.go' -or -name 'go.*' -not -path "./vendor/*")

bin/journal-cli: go.mod go.sum $(src)
	@mkdir -p bin
	@go build -o bin/journal-cli ./cmd/
