package errno

import (
	"github.com/huerni/gmitex/core/errno"
)

var (
	ServiceErr = errno.NewErrNo(500, "An error occurred within the service")
)
