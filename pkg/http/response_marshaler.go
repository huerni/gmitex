package http

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"io"
	"strings"
)

// Custom Marshaler
type CustomMarshaler struct {
	M *runtime.JSONPb
}

type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (c *CustomMarshaler) Marshal(v interface{}) ([]byte, error) {
	ty := fmt.Sprintf("%T", v)
	// 判断是否是error回应
	if strings.Contains(ty, "Err") {
		return c.M.Marshal(v)
	}

	return c.M.Marshal(&Response{
		Code:    200,
		Message: "Success",
		Data:    v,
	})
}

func (c *CustomMarshaler) Unmarshal(data []byte, v interface{}) error {
	return c.M.Unmarshal(data, v)
}

func (c *CustomMarshaler) NewDecoder(r io.Reader) runtime.Decoder {
	return c.M.NewDecoder(r)
}
func (c *CustomMarshaler) NewEncoder(w io.Writer) runtime.Encoder {
	return c.M.NewEncoder(w)
}
func (c *CustomMarshaler) ContentType(v interface{}) string {
	return "application/json"
}
