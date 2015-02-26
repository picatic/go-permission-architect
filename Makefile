# PHONY targets that do not have files
.PHONY: all test get-deps

all: test

get-deps:
	go get github.com/smartystreets/goconvey/convey

test: 
	go test -v ./ ./models ./example