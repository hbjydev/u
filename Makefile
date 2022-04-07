PREFIX=/usr/bin

all: clean build

clean:
	rm -rf u

build:
	go build -o u main.go

install:
	cp u ${PREFIX}/u
