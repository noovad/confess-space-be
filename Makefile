run:
	air -c .air.toml

dev:
	go run main.go

clean:
	rm -rf tmp/
	go clean -cache

gen:
	wire gen ./api

install:
	go mod tidy
	go mod download