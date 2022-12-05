
.PHONY: create-env
create-env:
	cp -n env_sample .env

.PHONY: osint-build
osint-build: 
	docker build -t run-shodan -f cmd/run-shodan/Dockerfile .

.PHONY: build
build: osint-build
	docker-compose build

