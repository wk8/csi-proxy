package main

import (
	"github.com/kubernetes-csi/csi-proxy/internal/server"
	file_system "github.com/kubernetes-csi/csi-proxy/internal/server/file_system"
)

func main() {
	s := server.NewServer(apiGroups()...)
	if err := s.Start(nil); err != nil {
		panic(err)
	}
}

// apiGroups returns the list of enabled API groups.
func apiGroups() []server.APIGroup {
	return []server.APIGroup{
		&file_system.Server{},
	}
}
