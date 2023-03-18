default: out/example

clean:
	rm -rf out

GetTestify:
	go get github.com/stretchr/testify
	go install github.com/stretchr/testify

test: *.go
	go test ./...

out/example: implementation.go cmd/example/main.go
	mkdir -p out
	go build -o out/example ./cmd/example
