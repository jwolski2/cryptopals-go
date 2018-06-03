bin=cryptopals
sources=$(wilcard *.go)

all: build check

build: $(sources)
	go build -o $(bin) $(sources)

check: $(sources)
	go test $(sources)
