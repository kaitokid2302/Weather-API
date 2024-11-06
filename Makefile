.PHONY: update

update:
	go get -u ./...
	go mod tidy