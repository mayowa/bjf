export GOPATH?=/data/files/snippets/go
export GOBIN=$(GOPATH)/bin

all: tools

test:
	go test -v -race

coverage:
	go test -cover

lint:
	go tool vet .

fmt:
	gofmt -d -s .

bench:
	go test -bench=.

tools: $(GOBIN)/bjf

clean:
	rm -f $(GOBIN)/bjf

$(GOBIN)/bjf: bjf.go cmd/bjf/main.go
	go install -v github.com/xor-gate/go-bjf/cmd/bjf
