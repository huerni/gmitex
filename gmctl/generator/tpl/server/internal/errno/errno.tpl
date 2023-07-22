package errno

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	{{.imports}}
)

type ErrNo struct {
	ErrCode codes.Code
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%v, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code {{.servicePackage}}.ErrCode, msg string) ErrNo {
	return ErrNo{
		ErrCode: codes.Code(code),
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo({{.servicePackage}}.ErrCode_SuccessCode, "Success")
	ServiceErr             = NewErrNo({{.servicePackage}}.ErrCode_ServiceErrCode, "An error occurred within the service")
	ParamErr               = NewErrNo({{.servicePackage}}.ErrCode_ParamErrCode, "Wrong Parameter has been given")

)

// ConvertErr convert error to Errno
func ConvertErr(err error) error {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return status.Error(Err.ErrCode, Err.ErrMsg)
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return status.Error(s.ErrCode, s.ErrMsg)
}

func ConvertErrNo(errno ErrNo) error {
	return status.Error(errno.ErrCode, errno.ErrMsg)
}
