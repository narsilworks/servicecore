package servicecore

import (
	"net/http"
	"time"

	ln "github.com/narsilworks/livenote"
)

type ServiceCore struct {
	// Configurable at service creation
	// This are manually coded at the derivative code
	id          string // A convenient ID for the service
	name        string // Name of the service
	copyright   string // Copyright information of the service
	description string // Description of the service
	version     string // Version of the service

	// Set before service starts
	//
	setter serviceSetter

	DefaultContentType string // Default content type of the service
	DefaultHostPort    int    //
	Production         bool   // Production mode flag
	ShutdownEvent      func() // Shutdown event processing for graceful exit

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

func Create(identity map[string]any) (*ServiceCore, error) {

	id := "SAMPLE"
	mapValue(&identity, "id", &id)

	return &ServiceCore{
		messages: *ln.NewLiveNote(id),
	}, nil
}

func (sc *ServiceCore) ID() string {
	return sc.id
}

func (sc *ServiceCore) Name() string {
	return sc.name
}

func (sc *ServiceCore) Description() string {
	return sc.description
}

func (sc *ServiceCore) Copyright() string {
	return sc.copyright
}

func (sc *ServiceCore) Version() string {
	return sc.version
}

func (sc *ServiceCore) Started() *time.Time {
	return &sc.started
}

func (sc *ServiceCore) Set() *serviceSetter {
	return &sc.setter
}

func (sc *ServiceCore) Get() *serviceGetter {
	return &serviceGetter{
		&sc.setter,
	}
}

func mapValue[T any](identity *map[string]any, key string, out *T) {
	if identity == nil || len(*identity) == 0 {
		return
	}

	val, ok := (*identity)["id"]
	if !ok {
		return
	}

	if rval, ok := val.(T); ok {
		*out = rval
	}
}

func Serve() {

}
