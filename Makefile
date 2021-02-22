.DEFAULT_GOAL := build

build:
	go build -o bin/webhook-plex cmd/main.go

deps:
	go mod vendor

clean:
	rm -r bin

test:
	go test -timeout 30s -count=1 ./plugins/inputs/webhooks/plex

run:
	./bin/webhook-plex --config plugin.conf
