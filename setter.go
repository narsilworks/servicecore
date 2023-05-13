package servicecore

import "github.com/narsilworks/servicecore/ifcs"

type serviceSetter struct {
	cache     ifcs.ICache         // Caching interface
	logger    ifcs.ILogger        // Logging interface. Required.
	queue     ifcs.IQueue         // Queue/Messaging interface
	data      []ifcs.IData        // Data access interface
	localData ifcs.ILocalData     // Local database provider
	cors      ifcs.ICORS          // CORS interface
	router    ifcs.IRouter        // Built-in router of the service. Required.
	config    ifcs.IConfiguration // Configuration settings of the service. Required.
}

// Config sets the available configuration for the service.
func (s *serviceSetter) Config(i ifcs.IConfiguration) {
	s.config = i
}

// Logger sets the logging mechanism of the server.
func (s *serviceSetter) Logger(i ifcs.ILogger) {
	s.logger = i
}

// Router sets the router of the service.
func (s *serviceSetter) Router(i ifcs.IRouter) {
	s.router = i
}

// Cache sets the cache mechanism
func (s *serviceSetter) Cache(i ifcs.ICache) {
	s.cache = i
}

// Queue sets the queue mechanism
func (s *serviceSetter) Queue(i ifcs.IQueue) {
	s.queue = i
}

// Data sets the data mechanism
func (s *serviceSetter) Data(i ...ifcs.IData) {
	s.data = i
}

// LocalData sets the local data mechanism
func (s *serviceSetter) LocalData(i ifcs.ILocalData) {
	s.localData = i
}

// CORS sets the cross-origin resource sharing of the service
func (s *serviceSetter) CORS(i ifcs.ICORS) {
	s.cors = i
}
