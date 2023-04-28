package servicecore

import (
	"net/http"
	"time"

	ln "github.com/narsilworks/livenote"
	"github.com/narsilworks/servicecore/dumm"
	"github.com/narsilworks/servicecore/ifcs"
)

type ServiceCore struct {
	ID                 string // A convenient ID for the service
	Name               string // Name of the service
	Copyright          string // Copyright information of the service
	Description        string // Description of the service
	DefaultContentType string // Default content type of the service
	DefaultHostPort    int    //
	Production         bool   // Production mode flag
	Version            string // Version of the service

	cache     ifcs.ICache         // Caching interface
	logger    ifcs.ILogger        // Logging interface. Required.
	queue     ifcs.IQueue         // Queue/Messaging interface
	data      []ifcs.IData        // Data access interface
	localData ifcs.ILocalData     // Local database provider
	cors      ifcs.ICORS          // CORS interface
	router    ifcs.IRouter        // Built-in router of the service. Required.
	config    ifcs.IConfiguration // Configuration settings of the service. Required.

	ShutdownEvent func() // Shutdown event processing for graceful exit

	appControllerURL string //

	httpServer *http.Server   //
	launchArgs map[string]any // Arguments upon launch
	messages   ln.LiveNote    // Message logged while launching the service

	mime map[string]string

	appInstance     string    //
	modInstance     string    // Module instance id
	publishedEvents []string  // List of the public events that the service publishes
	started         time.Time //
}

func Create() (*ServiceCore, error) {
	return &ServiceCore{}, nil
}

func (sc *ServiceCore) Cache() ifcs.ICache {
	if sc.cache == nil {
		return &dumm.Cache{}
	}
	return sc.cache
}

func (sc *ServiceCore) Logger() ifcs.ILogger {
	return sc.logger
}

func (sc *ServiceCore) Queue() ifcs.IQueue {
	return sc.queue
}

func (sc *ServiceCore) Data() []ifcs.IData {
	return sc.data
}

func (sc *ServiceCore) LocalData() ifcs.ILocalData {
	return sc.localData
}

func (sc *ServiceCore) CORS() ifcs.ICORS {
	if sc.cors == nil {
		return &dumm.CORS{}
	}
	return sc.cors
}

func (sc *ServiceCore) Router() ifcs.IRouter {
	return sc.router
}

func (sc *ServiceCore) Configuration() ifcs.IConfiguration {
	return sc.config
}
