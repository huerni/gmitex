version: '3'

services:
  mysql:
    image: mysql:latest
    restart: always
    container_name: {{.project}}-mysql
    environment:
      - MYSQL_DATABASE={{.project}}
      - MYSQL_USER={{.project}}
      - MYSQL_PASSWORD=123456
      - MYSQL_RANDOM_ROOT_PASSWORD="yes"
      - TZ = Asia/Shanghai
    ports:
      - 3306:3306

  Etcd:
    image: 'bitnami/etcd:latest'
    restart: always
    container_name: {{.project}}-etcd
    environment:
      - "ALLOW_NONE_AUTHENTICATION=yes"
      - "ETCD_ADVERTISE_CLIENT_URLS=http://etcd:2379"
    ports:
      - "2379:2379"
      - "2380:2380"

