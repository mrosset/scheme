.PHONY: check example repl

export GUILE_AUTO_COMPILE=0

default: check

repl:
	$(MAKE) -C $@

example:
	$(MAKE) check -C $@

check: repl example
	go test -v .

clean:
	-go clean ./repl
	-go clean ./example
