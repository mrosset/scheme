.PHONY: repl

export GUILE_AUTO_COMPILE=0

default: repl test

repl:
	$(MAKE) -C $@

check:
	go test -v ./pkg/...
