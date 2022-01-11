all:
	mkdir -p ./dist/bin/

	go build -ldflags="-w -s" -o ./dist/bin/wallblog ./cmd/wallblog.go
	chmod +x ./dist/bin/wallblog

fmt:
	gofmt -w ./cmd/*.go
	gofmt -w ./internal/cfg/*.go
	gofmt -w ./internal/handler/*.go

clean:
	rm -rf ./dist/
