prefix = "{{.project}}"

[etcd]
hosts = ["127.0.0.1:2379"]
key = "gateway"

[mysql]
username = "{{.project}}"
password = ""
address = "localhost:3306"
database = "{{.project}}"
other = "?charset=utf8&parseTime=True&loc=Local"

dsn = ""


