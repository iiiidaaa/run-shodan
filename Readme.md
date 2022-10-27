# Usage
## Build
```shell
docker build -t run-shodan -f cmd/run-shodan/Dockerfile .
```

## Run
```shell
docker run -e SHODAN_KEY={YOUR_API_KEY} run-shodan {target}
```