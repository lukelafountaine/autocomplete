run:
	docker run -p 80:80 autocomplete

build:
	GOOS=linux go build server.go trie.go
	docker build . -t autocomplete
