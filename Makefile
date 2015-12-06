export GOPATH?=/data/files/snippets/go
export GOBIN=$(GOPATH)/bin

all: tools

test:
	go test -v -race

lint:
	go tool vet .

fmt:
	gofmt -d -s .

bench:
	go test -bench=.

tools: $(GOBIN)/bjf

clean:
	rm -f $(GOBIN)/bjf

$(GOBIN)/bjf: bjf.go src/bjf.go
	go install src/bjf.go
