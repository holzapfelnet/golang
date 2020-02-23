# Go Docker Hello World

Rest API in Go using Docker 


## Prerequisites

* [Go](https://golang.org/dl/)	
* [Docker](https://docs.docker.com/docker-for-mac/)					
* [Visual Studio Code](https://code.visualstudio.com)			
  * [Go Extension for VS Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
  * [Docker Extension for VS Code](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker)

## Commands

go get github.com/stretchr/testify

go build

go test ./...

docker build -t go-docker-test .

docker run -d --name go-docker-container -p 8080:8080 go-docker-test

curl -i localhost:8080/greet

## Links

* [Writing Web Applications](https://golang.org/doc/articles/wiki)
* [Building Minimal Docker Containers for Go Applications](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications)