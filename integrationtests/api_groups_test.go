package integrationtests

import (
	"context"
	"testing"
	"time"

	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1alpha1"
	v1alpha1client "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/client/dummy/v1alpha1"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy"
	"github.com/kubernetes-csi/csi-proxy/internal/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIGroups(t *testing.T) {
	defer startServer(t)()

	t.Run("with v1alpha1", func(t *testing.T) {
		client, err := v1alpha1client.NewClient()
		require.Nil(t, err)
		defer client.Close()

		// happy path
		request := &v1alpha1.ComputeDoubleRequest{
			Input32: 28,
		}
		response, err := client.ComputeDouble(context.Background(), request)
		if assert.Nil(t, err) {
			assert.Equal(t, "wkpo", response.Response32)
		}
	})
}

// startServer starts the proxy's GRPC servers, and returns a function to shut them down when done with testing
func startServer(t *testing.T) func() {
	s := server.NewServer(&dummy.Server{})

	listeningChan := make(chan interface{})
	go func() {
		assert.Nil(t, s.Start(listeningChan))
	}()

	select {
	case <-listeningChan:
	case <-time.After(5 * time.Second):
		t.Fatalf("Timed out waiting for GRPC servers to start listening")
	}

	return func() {
		assert.Nil(t, s.Stop())
	}
}
