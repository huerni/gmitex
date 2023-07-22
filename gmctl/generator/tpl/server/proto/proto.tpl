syntax = "proto3";
option go_package = "./pb";
package {{.package}};

import "google/api/annotations.proto";

enum ErrCode {
  NormalCode                 = 0;
  SuccessCode                = 200;
  ServiceErrCode             = 500;
  ParamErrCode               = 10002;
}

message PingReq {
      string ping = 1;
}

message PingResp {
      string pong = 1;
}

service {{.serverName}} {
      rpc Ping(PingReq) returns (PingResp) {
            option (google.api.http) = {
                  get: "/api/v1.0/ping"
            };
      };
}