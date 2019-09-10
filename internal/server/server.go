package server

import (
	"fmt"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy"
	"github.com/kubernetes-csi/csi-proxy/internal"
	"net"
	"sync"

	"github.com/Microsoft/go-winio"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Server struct {
	versionedAPIs []*VersionedAPI
	started       bool
	mutex         *sync.Mutex
	grpcServers   []*grpc.Server
}

// TODO wkpo comment?
func NewServer(apiGroups ...APIGroup) *Server {
	if len(apiGroups) == 0 {
		apiGroups = defaultAPIGroups()
	}

	versionedAPIs := make([]*VersionedAPI, 0, len(apiGroups))
	for _, apiGroup := range apiGroups {
		versionedAPIs = append(versionedAPIs, apiGroup.VersionedAPIs()...)
	}

	return &Server{
		versionedAPIs: versionedAPIs,
		mutex:         &sync.Mutex{},
	}
}

// TODO wkpo circular import la non? this should be in a different pkg, prolly - maybe a config pkg? together with the named pipes' prefix?
// TODO wkpo comment
func defaultAPIGroups() []APIGroup {
	// TODO: add API groups as we add them to the project
	// TODO wkpo
	return []APIGroup{&dummy.Server{}}
}

// Starts starts one GRPC server per API version; it is a blocking call, that returns
// as soon as any of those servers shuts down (at which point it also shuts down all the
// others).
// If passed a listeningChan, it will close it when it's started listening.
// TODO wkpo chan to say when started listening? use for tests!!
func (s *Server) Start(listeningChan chan interface{}) []error {
	doneChan, errors := s.startListening()
	if len(errors) != 0 {
		return errors
	}
	defer close(doneChan)

	if listeningChan != nil {
		close(listeningChan)
	}

	return s.waitForGRPCServersToStop(doneChan)
}

// startListening creates the named pipes, and starts GRPC servers listening on them.
func (s *Server) startListening() (chan *versionedAPIDone, []error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.started {
		return nil, []error{fmt.Errorf("server already started")}
	}
	s.started = true

	listeners, errors := s.createListeners()
	if len(errors) != 0 {
		return nil, errors
	}

	return s.createAndStartGRPCServers(listeners), nil
}

// createListeners creates the named pipes.
func (s *Server) createListeners() (listeners []net.Listener, errors []error) {
	listeners = make([]net.Listener, len(s.versionedAPIs))

	for i, versionedAPI := range s.versionedAPIs {
		pipePath := internal.PipePath(versionedAPI.Group, versionedAPI.Version)

		listener, err := winio.ListenPipe(pipePath, nil)
		if err == nil {
			listeners[i] = listener
		} else {
			errors = append(errors, err)
		}
	}

	if len(errors) != 0 {
		// let's do a best effort to close all the listeners that we did manage to create
		for _, listener := range listeners {
			if listener != nil {
				listener.Close()
			}
		}
	}

	return
}

type versionedAPIDone struct {
	i   int
	err error
}

// createAndStartGRPCServers creates the GRPC servers, but doesn't start them just yet.
func (s *Server) createAndStartGRPCServers(listeners []net.Listener) chan *versionedAPIDone {
	doneChan := make(chan *versionedAPIDone, len(listeners))

	for i, versionedAPI := range s.versionedAPIs {
		grpcServer := grpc.NewServer()
		s.grpcServers[i] = grpcServer

		versionedAPI.Registrant(grpcServer)

		go func() {
			err := grpcServer.Serve(listeners[i])

			doneChan <- &versionedAPIDone{
				i:   i,
				err: err,
			}
		}()
	}

	return doneChan
}

func (s *Server) waitForGRPCServersToStop(doneChan chan *versionedAPIDone) (errs []error) {
	processServerDoneEvent := func(event *versionedAPIDone) {
		if event.err != nil {
			versionedAPI := s.versionedAPIs[event.i]
			err := errors.Wrapf(event.err, "GRPC server for API group %s version %s failed", versionedAPI.Group, versionedAPI.Version)
			errs = append(errs, err)
		}
	}

	// and now let's wait for at least one server to be done
	processServerDoneEvent(<-doneChan)

	// let's stop all other servers
	s.Stop()

	// and wait for them to stop
	// TODO: do we want a timeout here?
	for doneCount := 1; doneCount < len(s.versionedAPIs); doneCount++ {
		processServerDoneEvent(<-doneChan)
	}

	return
}

// Stop stops all GRPC servers.
func (s *Server) Stop() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.started {
		return fmt.Errorf("server not started yet")
	}

	for _, grpcServer := range s.grpcServers {
		if grpcServer != nil {
			grpcServer.Stop()
		}
	}

	return nil
}
