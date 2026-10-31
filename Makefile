run:
	DEBUG=1 go run .

logs:
	tail -f debug.log

build:
	@mkdir -p ./dist
	@CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o ./dist/sweeper-tui .
	@echo "program built to ./dist dir"

clean:
	rm -rf ./dist

.PHONY: run build clean