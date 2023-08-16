package errno

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrNo struct {
	ErrCode codes.Code
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%v, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{
		ErrCode: codes.Code(code),
		ErrMsg:  msg,
	}
}

func ConvertErrNo(errno ErrNo) error {
	return status.Error(errno.ErrCode, errno.ErrMsg)
}

func Err(code int, msg string) error {
	err := ErrNo{
		ErrCode: codes.Code(code),
		ErrMsg:  msg,
	}

	return status.Error(err.ErrCode, err.ErrMsg)
}

func ConvertErr(err error) error {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return status.Error(Err.ErrCode, Err.ErrMsg)
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return status.Error(s.ErrCode, s.ErrMsg)
}

var (
	Success    = NewErrNo(200, "Success")
	ServiceErr = NewErrNo(500, "An error occurred within the service")
	ParamErr   = NewErrNo(503, "Wrong Parameter has been given")
)
