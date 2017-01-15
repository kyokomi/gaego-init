APP_ROOT=appengine
VENDOR_GOPATH=$(shell pwd)/$(APP_ROOT)/gopath/vendor
APP_GOPATH=$(shell pwd)/$(APP_ROOT)/gopath
RUN_GOPATH=$(APP_GOPATH):$(VENDOR_GOPATH)

install:
	glide init
	glide install
	mkdir -p $(APP_ROOT)/gopath/vendor
	ln -s $(shell pwd)/vendor/ $(VENDOR_GOPATH)/src

run:
	GOPATH=$(RUN_GOPATH) goapp serve $(APP_ROOT)

clean-run:
	GOPATH=$(RUN_GOPATH) goapp serve --clear_datastore $(APP_ROOT)

deploy:
	GOPATH=$(RUN_GOPATH) goapp deploy $(APP_ROOT)

