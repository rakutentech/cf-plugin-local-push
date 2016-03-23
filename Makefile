VERSION = 0.1.0
COMMIT = $$(git describe --always)

default: build

deps:
	go get -v .

build: deps
	go build -ldflags "-X main.GitCommit=$(COMMIT)" -o bin/cf-plugin-local-push

install: build
	cf install-plugin bin/cf-plugin-local-push -f
	cf plugins

xbuild: deps
	@if [ -d "out/$(VERSION)" ]; then rm -fr out; fi
	gox \
      -ldflags "-X main.GitCommit=$(COMMIT)" \
      -parallel=3 \
      -os="darwin linux windows" \
      -arch="amd64" \
      -output "out/$(VERSION)/{{.Dir}}_{{.OS}}_{{.Arch}}"
	cd out/$(VERSION) && shasum * > SHASUMS && cat SHASUMS

release: 
	ghr $(VERSION) out/$(VERSION)

uninstall:
	cf uninstall-plugin 'local-push'

test: vet 
	go test -v

vet:
	@go get golang.org/x/tools/cmd/vet
	go tool vet *.go

lint:
	@go get github.com/golang/lint/golint
	golint ./...

# cover shows test coverages
cover:
	@go get golang.org/x/tools/cmd/cover		
	godep go test -coverprofile=cover.out
	go tool cover -html cover.out
	rm cover.out
