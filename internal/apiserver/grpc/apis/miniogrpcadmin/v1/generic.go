package v1

import (
	"github.com/juju/loggo"

	"github.com/venezia/minio-grpc-admin/pkg/util/log"
)

var (
	logger loggo.Logger
)

type Server struct{}

func init() {
	logger = log.GetModuleLogger("internal.apiserver.grpc.apis.miniogrpcadmin.v1")
}
