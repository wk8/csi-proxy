package server

import (
	"fmt"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy"
	"net"
	"sync"

	"github.com/Microsoft/go-winio"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Server struct {
	apiGroupServers []ApiGroupServer
	started         bool
	mutex           *sync.Mutex
	grpcServers     []*grpc.Server
}

// TODO wkpo comment?
func NewServer(apiGroups ...ApiGroupServer) *Server {
	if len(apiGroups) == 0 {
		apiGroups = defaultApiGroups()
	}

	return &Server{
		apiGroupServers: apiGroups,
		mutex:           &sync.Mutex{},
	}
}

// TODO wkpo circular import la non?
// TODO wkpo comment
func defaultApiGroups() []ApiGroupServer {
	// TODO: add API groups as we add them to the project
	// TODO wkpo
	return []ApiGroupServer{&dummy.Server{}}
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
func (s *Server) startListening() (chan *apiVersionServerDone, []error) {
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

// createListeners creates the named pipes
func (s *Server) createListeners() (listeners []net.Listener, errors []error) {
	definitions := s.handler.Definitions()
	listeners = make([]net.Listener, len(definitions))

	for i, definition := range definitions {
		pipePath := internal.PipePathForApiVersion(definition.Version)

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

// createAndStartGRPCServers creates the GRPC servers, but doesn't start them yet.
func (s *Server) createAndStartGRPCServers(listeners []net.Listener) chan *apiVersionServerDone {
	definitions := s.handler.Definitions()
	s.grpcServers = make([]*grpc.Server, len(definitions))

	doneChan := make(chan *apiVersionServerDone, len(s.handler.Definitions()))

	for i, definition := range definitions {
		grpcServer := grpc.NewServer()
		s.grpcServers[i] = grpcServer

		definition.BuildAndRegisterServers(grpcServer, definition.Version, s.handler)

		go func() {
			err := grpcServer.Serve(listeners[i])

			doneChan <- &apiVersionServerDone{
				definitionIndex: i,
				err:             err,
			}
		}()
	}

	return doneChan
}

// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies
// TODO wkpo oldies

type Server struct {
	handler     *apiversions.Handler
	started     bool
	mutex       *sync.Mutex
	grpcServers []*grpc.Server
}

// TODO wkpo comment?
func NewServer(definitions ...apiversions.Definition) (*Server, error) {
	handler := &apiversions.Handler{}
	if err := handler.Register(definitions...); err != nil {
		return nil, err
	}

	return &Server{
		handler: handler,
		mutex:   &sync.Mutex{},
	}, nil
}

type apiVersionServerDone struct {
	definitionIndex int
	err             error
}

func (s *Server) waitForGRPCServersToStop(doneChan chan *apiVersionServerDone) (errs []error) {
	definitions := s.handler.Definitions()

	processServerDoneEvent := func(event *apiVersionServerDone) {
		if event.err != nil {
			err := errors.Wrapf(event.err, "GRPC server for API version %v failed", definitions[event.definitionIndex].Version)
			errs = append(errs, err)
		}
	}

	// and now let's wait for at least one server to be done
	processServerDoneEvent(<-doneChan)

	// let's stop all other servers
	s.Stop()

	// and wait for them to stop
	// TODO: do we want a timeout here?
	for doneCount := 1; doneCount < len(definitions); doneCount++ {
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

func (s *Server) WrapHandler() {
	// TODO wkpo next ?
}
