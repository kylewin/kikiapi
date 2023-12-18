build_date := ${shell date -u "+%Y-%m-%d%H:%M:%SUTC"}
build_version := 1.0.0

build:
	go build -ldflags "-X main.buildVersion=$(build_version) -X main.buildDate=$(build_date)"

run: build
	./kikiapi

.PHONY: build
