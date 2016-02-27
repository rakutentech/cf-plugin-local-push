default: build

build:
	go build -o bin/cf-plugin-local-push

install: build
	cf install-plugin bin/cf-plugin-local-push -f
	cf plugins

uninstall:
	cf uninstall-plugin 'local-push'