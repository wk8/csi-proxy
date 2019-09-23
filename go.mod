module github.com/kubernetes-csi/csi-proxy

go 1.12

replace github.com/kubernetes-csi/csi-proxy/client => ./client

require (
	github.com/Microsoft/go-winio v0.4.14
	github.com/golang/protobuf v1.3.2
	github.com/kubernetes-csi/csi-proxy/client v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.8.1
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.4.0
	google.golang.org/grpc v1.23.1
	k8s.io/gengo v0.0.0-20190907103519-ebc107f98eab
	k8s.io/klog v1.0.0 // indirect
)
