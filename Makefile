init:
	go get -u github.com/jteeuwen/go-bindata/...
build:
	go generate
	go build
install:
	go generate
	go install

