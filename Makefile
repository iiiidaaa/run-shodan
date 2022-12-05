PHONY: build
build: 
	create-env
	osint-build
	docker-compose build
osint-build: 
	docker build -t run-shodan -f cmd/run-shodan/Dockerfile .
create-env:
	cp -n env_sample .env
