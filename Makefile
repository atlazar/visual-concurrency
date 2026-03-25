cli:
	go build -o ./bin/ ./cmd/cli/
	./bin/cli

gui:
	go build -o ./bin/ ./cmd/gui/
	./bin/gui