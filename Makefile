PHONY: build
build: 
	docker build -t run-shodan -f cmd/run-shodan/Dockerfile .