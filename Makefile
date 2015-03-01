prefix=/usr/local

ws:
	go build ws.go

all: ws

install: all
	cp ws $(prefix)
	cp ws.1 $(prefix)/share/man/man1/

clean:
	rm -rf *~ ws
