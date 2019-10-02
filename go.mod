module github.com/kubernetes-csi/csi-proxy

go 1.12

require (
	github.com/Microsoft/go-winio v0.4.14
	github.com/golang/protobuf v1.3.2
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365
	github.com/kubernetes-csi/csi-proxy/client v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.4.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.4.0
	github.com/wk8/go-ordered-map v0.1.0
	google.golang.org/grpc v1.23.1
	k8s.io/gengo v0.0.0-20190907103519-ebc107f98eab
	k8s.io/klog v1.0.0
)

replace (
	github.com/kubernetes-csi/csi-proxy/client => ./client
	// using my fork of gengo until
	// https://github.com/kubernetes/gengo/pull/155#issuecomment-537589085
	// is implemented, and the generic conversion generator merged into code-generator
	k8s.io/gengo => github.com/wk8/gengo v0.0.0-20191001015530-2c8d7500e6ee7d247f53ce7f3bbc91a7105291ec
)
