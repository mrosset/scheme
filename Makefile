.PHONY: repl

export GUILE_AUTO_COMPILE=0

default: repl test

repl:
	$(MAKE) -C $@

test:
	go clean -cache
	go test -v ./pkg/...
