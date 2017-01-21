run:
	docker run -p 80:80 autocomplete

build:
	GOOS=linux go build src/server.go src/trie.go
	docker build . -t autocomplete
