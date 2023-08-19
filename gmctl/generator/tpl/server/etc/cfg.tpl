prefix = "{{.projectName}}"

[grpc]
name = "{{.serverName}}"
listenOn = "127.0.0.1:8081"

[http]
listenOn = "127.0.0.1:8080"

[etcd]
hosts = ["127.0.0.1:2379"]
key = "{{.serverName}}"

[mysql]
# etcd 默认mysql
key = "mysql"

# 自定义
username = ""
password = ""
address = ""
database = ""
other = ""

dsn = ""

[traefik]
provider = "etcd"