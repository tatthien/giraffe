.PHONY: clean build

clean: 
	rm -rf ./bin

build: clean
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -o  bin/giraffe.linux-amd64 main.go 
	GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -a -o  bin/giraffe.darwin-amd64 main.go 
	upx --brute bin/giraffe.linux-amd64
	upx --brute bin/giraffe.darwin-amd64

# Example: make release V=0.0.0
release:
	git tag v$(V)
	@read -p "Please enter to confirm and push to origin ..." && git push origin v$(V)