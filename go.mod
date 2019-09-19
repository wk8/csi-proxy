module github.com/wk8/csi-proxy

go 1.12

require (
	github.com/Microsoft/go-winio v0.4.14
	github.com/golang/protobuf v1.3.2
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.4.0
	github.com/wk8/csi-proxy/api v0.0.0-00010101000000-000000000000
	github.com/wk8/csi-proxy/client v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.23.1
	k8s.io/api v0.0.0-20190919035539-41700d9d0c5b
)

replace github.com/wk8/csi-proxy/api => ./api

replace github.com/wk8/csi-proxy/client => ./client
