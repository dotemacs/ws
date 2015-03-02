PREFIX=/usr/local

ws:
	go build ws.go

all: ws

install: all
	cp ws $(PREFIX)
	cp ws.1 $(PREFIX)/share/man/man1/

clean:
	rm -rf *~ ws
