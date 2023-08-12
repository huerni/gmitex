package errno

import (
	"github.com/huerni/gmitex/pkg/errno"
	{{.imports}}
)


var (
	ServiceErr = errno.NewErrNo(500, "An error occurred within the service")
)
