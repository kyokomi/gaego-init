init:
	go get -u github.com/jteeuwen/go-bindata/...
build:
	go generate
	go build
install: init
	go generate
	go install

