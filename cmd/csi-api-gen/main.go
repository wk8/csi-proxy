package main

import (
	"os"

	"github.com/kubernetes-csi/csi-proxy/cmd/csi-api-gen/execute"
)

func main() {
	execute.Execute(os.Args[1:])
	// TODO wkpo should succeed, then remove!
	execute.Execute(os.Args[1:])
}
