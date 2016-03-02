# cf-plugin-local-push

`cf-plugin-local-push` is a [cloudfoundry/cli](https://github.com/cloudfoundry/cli) plugin. It allows you to push your cloudfoundry application to your local docker container with actual [buildpacks](http://docs.cloudfoundry.org/buildpacks/) :whale:. This plugin manipulates [DEA](https://docs.cloudfoundry.org/concepts/architecture/execution-agent.html) (where cf application is runnging) enviroment. So this can be used for setting up very light weight debug environment for application developers. And power of docker build cache, start up application is really *fast*.

This plugin is still *PoC*, so please be careful to use this plugin.  

## Why?

Why we need this? Because the application developers (at least, me) want to debug their cf app on local environment before `push`-ing to actual environment. Since it's faster and you don't need care about breaking the app or wasting resources (you may not have internet access when they need to run it), it's important to have local development environment.

Cloudfoundry community provides [bosh-lite](https://github.com/cloudfoundry/bosh-lite) for local dev environment for BOSH using warden containers. But for me, it's too heavy and not for **user**. It's only for CF operators. 

## Demo

The following is demo to run sample ruby application (the code is available [here](/sample)).

![demo]()

Just `cf local-push`, it detects application runtime and start building application with its buildpack. While it takes time at first push, it's really fast in the second time because of docker build cache.

## Install

To install this plugin, use `go get` (make sure you have already setup golang enviroment like `$GOPATH`),

```bash
$ go get -d github.com/tcnksm/cf-plugin-local-push
$ cd $GOPATH/src/github.com/tcnksm/cf-plugin-local-push
$ make install # if you have already installed, then run `make uninstall` before
```

Since this plugin is still immature and PoC, it's not uploaded on [Community Plugin Repo](http://plugins.cloudfoundry.org/ui/). But in future, I'll add this plugin there and make it more easy to install.

## Usage

To use this plugin, you need to setup docker environment, docker daemon running and docker client cli (See [Docker Toolbox](https://www.docker.com/products/docker-toolbox)). Then run the following command in the directory where your application source is.

```bash
$ cf local-push
```

**NOTE1**: This plugins does not support parsing `manifest.yml` yet. Currently, it's only manipulate executing buildpack and parsing `Procfile`.

**NOTE2**: Currently it uses [gliderlabs/herokuish](https://github.com/gliderlabs/herokuish) inside base image, so buildpack is heroku's one. So it' a bit different from cf buildpack. It will be replaced with CF buildpack.

**NOTE3**: It's not allowed to use arbittrary buildpack now. Check the available buildpack [here](https://github.com/gliderlabs/herokuish/tree/master/buildpacks).

`local-push` will a build docker image with compiling your application source code by appropriate buildpack. After building, you can access to an application runnging (by default, port is `8080`),

```bash
$ curl $(docker-machine ip):8080
```

## Contribution

1. Fork ([https://github.com/tcnksm/cf-plugin-local-push/fork](https://github.com/tcnksm/cf-plugin-local-push/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[Taichi Nakashima](https://github.com/tcnksm)
