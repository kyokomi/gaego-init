APP_ROOT=appengine
VENDOR_GOPATH=$(shell pwd)/$(APP_ROOT)/gopath/vendor
APP_GOPATH=$(shell pwd)/$(APP_ROOT)/gopath
RUN_GOPATH=$(APP_GOPATH):$(VENDOR_GOPATH)

install:
	go get -u github.com/golang/dep/...
	mkdir -p $(APP_ROOT)/gopath/vendor
	dep init
	dep ensure -v -update
	ln -s $(shell pwd)/vendor/ $(VENDOR_GOPATH)/src

dep:
	dep ensure -v

run: dep
	GOPATH=$(RUN_GOPATH) goapp serve $(APP_ROOT)

clean-run: dep
	GOPATH=$(RUN_GOPATH) goapp serve --clear_datastore $(APP_ROOT)

deploy: dep
	GOPATH=$(RUN_GOPATH) goapp deploy $(APP_ROOT)

