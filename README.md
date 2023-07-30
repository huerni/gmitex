# gmitex
A Go microservices framework  

gmitex is a lightweight microservices framework designed for individuals to quickly learn and develop microservices. Each microservice is configured with both RPC and RESTful interfaces and exposed to the outside world through a gateway. It supports dynamic addition of routes without the need for redundant definition of API files.

**The purpose of this project is for personal learning and building purposes. Currently, it is at the demo level, with plenty of areas that can be modified and improved.**

## Env
go 1.20  
docker 24.0  

# Install
```shell
go install github.com/huerni/gmitex/gmctl@latest
```

## Quick Start
```shell
#env check
gmctl check

# Code generation
gmctl gateway
gmctl new <serverName>

#cd gateway
docker-compose up -d
go mod tidy
go run cmd/main.go

#cd server
go mod tidy
go run cmd/main.go
```

