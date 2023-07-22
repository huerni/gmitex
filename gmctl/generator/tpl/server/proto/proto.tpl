syntax = "proto3";
option go_package = "./pb";
package {{.package}};

import "google/api/annotations.proto";

message PingReq {
      string ping = 1;
}

message PingResp {
      string pong = 1;
}

service {{.package}} {
      rpc Ping(PingReq) returns (PingResp) {
            option (google.api.http) = {
                  get: "/api/v1.0/ping"
            };
      };
}