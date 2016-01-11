GOBL=$(GOPATH)/bin/gobl
GOBL_SRC=$(shell find . -name '*.go') 

$(GOBL): $(GOBL_SRC)
	go install github.com/toshaf/gobl

run: $(GOBL)
	gobl pack github.com/toshaf/exhibit
	tar -tf $(GOPATH)/gobl-pkg/github.com.toshaf.exhibit.gobl

runv: $(GOBL)
	gobl pack -v github.com/toshaf/exhibit
	tar -tf $(GOPATH)/gobl-pkg/github.com.toshaf.exhibit.gobl

test:
	go test github.com/toshaf/gobl/...

fmt:
	go fmt github.com/toshaf/gobl/...

e2e: $(GOBL)
	cd test-files/ && ./exec.sh
