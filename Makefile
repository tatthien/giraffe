.PHONY: clean build

clean: 
	rm -rf ./bin

build: clean
	GOOS=linux GOARCH=amd64 go build -o bin/giraffe.linux-amd64 main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/giraffe.darwin-amd64 main.go

release:
	git tag v$(V)
	@read -p "Please enter to confirm and push to origin ..." && git push origin v$(V)