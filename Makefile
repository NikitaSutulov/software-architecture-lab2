default: out/example

clean:
	rm -rf out

GetTestify:
	go get https://github.com/stretchr/testify
	go install https://github.com/stretchr/testify

test: *.go
	go test ./...

out/example: implementation.go cmd/example/main.go
	mkdir -p out
	go build -o out/example ./cmd/example
