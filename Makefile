build:
	cd src/knowsearch.ml && go build -o ../../jsonops

install:
	cd src/knowsearch.ml && go mod download

digger:
	make build && ./jsonops

validator:
	make build && ./jsonops --pretty=1 --verbose=1 check.txt

test:
	cd src/knowsearch.ml && go test ./...