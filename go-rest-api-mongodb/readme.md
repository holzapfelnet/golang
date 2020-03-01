# Simple API to store and receive todo items

Simple API to store and receive todo items. 
The todo items are stored within a MongoDB.
API and MongoDB are bundled within a docker-compose file.


## Prerequisites

* [Go](https://golang.org/dl/)	
* [Docker](https://docs.docker.com/docker-for-mac/)					
* [Visual Studio Code](https://code.visualstudio.com)			
  * [Go Extension for VS Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
  * [Docker Extension for VS Code](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker)

## Commands

docker-compose up --build

curl -i localhost:8081

curl -i -X POST -H "Content-Type: application/json" -d '{"Text":"New Item", "Done": false }' localhost:8081/add

## Links

* [Writing Web Applications](https://golang.org/doc/articles/wiki)
* [Building Minimal Docker Containers for Go Applications](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications)
* [Quickstart: Build a console app using Azure Cosmos DB's API for MongoDB and Golang SDK](https://docs.microsoft.com/en-us/azure/cosmos-db/create-mongodb-golang)
* [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)
* [MongoDB Go Driver Tutorial](https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial)