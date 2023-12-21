server:
    port: 8790
    contextPath: /
    timeout: 10
    host:
registerCenter:
    refreshFrequency: 30
#    eureka:
#        serviceUrls: [ http://localhost:8761 ]
    etcd:
        prefix: test
        endpoints: ["127.0.0.1:2379"]

gateway:
    routers:
        gmitest:
            path: /api/v1/gmitest/**
            serviceId: gmitest
            stripPrefix: false
            timeout: 30
