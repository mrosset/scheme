CFLAGS	= `pkg-config --cflags guile-2.2`
LDFLAGS	= `pkg-config --libs guile-2.2`

default: main

main: main.c
	gcc main.c -o main $(CFLAGS) $(LDFLAGS)

clean:
	-rm main
