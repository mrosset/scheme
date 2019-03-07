.PHONY: repl

default: repl test

repl:
	go build -o ./$@/$@ ./$@
	file ./$@/$@
	du -hs ./$@/$@

test:
	go test -v ./pkg/...
