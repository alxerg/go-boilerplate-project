# Makefile
#
build:
	gb build

test:
	gb test

re-test:
	reflex -r src/ $(MAKE) test

dev-deps:
	go get github.com/cespare/reflex
	go get github.com/constabulary/gb/...
