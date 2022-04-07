PREFIX=/usr/bin

all: clean build

clean:
	rm u

build:
	go build -o u main.go

install:
	cp u ${PREFIX}/u
