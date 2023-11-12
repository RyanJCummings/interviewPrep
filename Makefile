build:
	@go build -o bin/interviewPrep

run: build
	@./bin/interviewPrep

test:
	@go test -v ./...
