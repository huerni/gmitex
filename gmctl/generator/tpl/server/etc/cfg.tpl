prefix = "test"

[grpc]
name = "gmitest"
addr = "127.0.0.1"
port = 8081

[http]
addr = "127.0.0.1"
port = 8080

[etcd]
hosts = ["127.0.0.1:2379"]
key = "gmitest"


#[mysql]
## etcd 默认mysql
#key = "mysql"
#
## 自定义
#username = ""
#password = ""
#address = ""
#database = ""
#other = ""
#
#dsn = ""