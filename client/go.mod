module github.com/wk8/csi-proxy/client

go 1.12

require (
	github.com/Microsoft/go-winio v0.4.14
	github.com/pkg/errors v0.8.1
	github.com/wk8/csi-proxy/api v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.23.1
)

replace github.com/wk8/csi-proxy/api => ./../api
