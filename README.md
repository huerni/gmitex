<div align="center">
<h1>gmitex</h1>
</div>

<div align="center">
<b>A Go microservices framework</b>
</div>

<div align="center">
<img src="https://img.shields.io/badge/Golang-1.20-orange"/>
<img src="https://img.shields.io/badge/Etcd-latest-green"/>
<img src="https://img.shields.io/badge/MySQL-8.0-yellowgreen"/>
<img src="https://img.shields.io/badge/Docker-24.0.6-blue"/>
</div>

gmitex is a lightweight microservices framework designed for individuals to quickly learn and develop microservices. Each microservice is configured with both RPC and RESTful interfaces and exposed to the outside world through a gateway. It supports dynamic addition of routes without the need for redundant definition of API files.

>**The purpose of this project is for personal learning and building purposes. Currently, it is at the demo level, with plenty of areas that can be modified and improved.**

# Install
```shell
go install github.com/huerni/gmitex/gmctl@latest
```

## Quick Start
```shell
#env check
gmctl check
mkdir <ProjectDir>

# Code generation
gmctl gateway
gmctl new <ServerName>

#cd gateway
docker-compose up -d
go mod tidy
go run cmd/main.go

#cd server
go mod tidy
go run cmd/main.go
```


