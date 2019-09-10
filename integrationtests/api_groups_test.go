package integrationtests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1alpha1"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1alpha2"
	v1client "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/client/dummy/v1"
	v1alpha1client "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/client/dummy/v1alpha1"
	v1alpha2client "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/client/dummy/v1alpha2"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy"
	"github.com/kubernetes-csi/csi-proxy/internal/server"
)

func TestAPIGroups(t *testing.T) {
	defer startServer(t)()

	t.Run("with v1alpha1", func(t *testing.T) {
		client, err := v1alpha1client.NewClient()
		require.Nil(t, err)
		defer client.Close()

		request := &v1alpha1.ComputeDoubleRequest{
			Input32: 28,
		}
		response, err := client.ComputeDouble(context.Background(), request)
		if assert.Nil(t, err) {
			assert.Equal(t, 56, response.Response32)
		}
	})

	t.Run("with v1alpha2", func(t *testing.T) {
		client, err := v1alpha2client.NewClient()
		require.Nil(t, err)
		defer client.Close()

		request := &v1alpha2.ComputeDoubleRequest{
			Input64: 28,
		}
		response, err := client.ComputeDouble(context.Background(), request)
		if assert.Nil(t, err) {
			assert.Equal(t, int64(56), response.Response)
		}
	})

	t.Run("with v1", func(t *testing.T) {
		client, err := v1client.NewClient()
		require.Nil(t, err)
		defer client.Close()

		request := &v1.TellMeAPoemRequest{
			IWantATitle: true,
		}
		response, err := client.TellMeAPoem(context.Background(), request)
		if assert.Nil(t, err) {
			assert.Equal(t, "The New Colossus", response.Title)
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
