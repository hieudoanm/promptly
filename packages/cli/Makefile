lint:
	go vet ./...

format:
	go fmt ./...

build:
	go vet ./...
	go fmt ./...
	rm -rf ./bin
	go build -o ./bin/chat ./main.go

install:
	go vet ./...
	go fmt ./...
	rm -rf ./bin
	go build -o ./bin/chat ./main.go
	mkdir -p ~/bin
	cp ./bin/chat ~/bin/
	chmod +x ~/bin/chat
	export PATH="$(HOME)/bin:$(PATH)"
