# Usage
## Build
```shell
docker build -t run-shodan -f cmd/run-shodan/Dockerfile .
```

## Run
formatには、`json`,`text`を指定してください
```shell
docker run -e SHODAN_KEY={YOUR_API_KEY} run-shodan {format} {target}
```