GAE_GOPATH=$(shell pwd)/vendor
GAE_GOPATH2=$(shell pwd)

install:
	mkdir -p $(GAE_GOPATH)

run:
	GOPATH=$(GAE_GOPATH) goapp serve src/app.yaml

test:
	GOPATH=$(GAE_GOPATH):$(GAE_GOPATH2) goapp test -v ./src/...

deploy:
	GOPATH=$(GAE_GOPATH) goapp deploy src/app.yaml
