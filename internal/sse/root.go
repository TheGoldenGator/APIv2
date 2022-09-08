package sse

import (
	"context"
	"net/http"
	"sync"
)

type connection struct {
	writer     http.ResponseWriter
	flusher    http.Flusher
	requestCtx context.Context
}

// Server - instance used to accept connections and store them
type ServerT struct {
	connections      map[string]*connection
	connectionsMutex sync.RWMutex
}

var ServerTest *ServerT

func init() {
	ServerTest = new(ServerT)
}

func NewServer() *ServerT {
	server := &ServerT{
		connections: map[string]*connection{},
	}

	return server
}

func (s *ServerT) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	flusher, ok := rw.(http.Flusher)

	if !ok {
		http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	requestContext := req.Context()
	s.connectionsMutex.Lock()
	s.connections[req.RemoteAddr] = &connection{
		writer:     rw,
		flusher:    flusher,
		requestCtx: requestContext,
	}
	s.connectionsMutex.Unlock()

	defer func() {
		s.removeConnection(req.RemoteAddr)
	}()

	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	<-requestContext.Done()
}

func (s *ServerT) removeConnection(client string) {
	s.connectionsMutex.Lock()
	defer s.connectionsMutex.Unlock()

	delete(s.connections, client)
}
