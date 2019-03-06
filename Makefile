SRC	= scheme.x $(wildcard *.go) Makefile
CFLAGS	= `pkg-config --cflags --libs guile-2.2`

export CGO_ENABLED=1

scheme.x: scheme.h
	guile-snarf -o $@ $< $(CFLAGS)
