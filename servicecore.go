package servicecore

import (
	"fmt"
	"net/http"
	"time"

	ln "github.com/narsilworks/livenote"
	"github.com/narsilworks/servicecore/ifcs"
	"github.com/segmentio/ksuid"
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
	setter serviceSetter

	DefaultContentType string // Default content type of the service
	DefaultHostPort    int    //
	OnShutdown         func() // Shutdown event processing for graceful exit

	appControllerURL string //

	httpServer *http.Server   //
	launchArgs map[string]any // Arguments upon launch
	messages   ln.LiveNote    // Message logged while launching the service

	appInstance     string    //
	modInstance     string    // Module instance id
	publishedEvents []string  // List of the public events that the service publishes
	started         time.Time //
	productionMode  bool      // Indicates if the service is in production mode or not
	hostPort        int       // the port where the service listens

}

func Create(identity map[string]any) (*ServiceCore, error) {

	id := "SAMPLE"
	name := "Sample Service"
	copyright := "Copyright 2023, NarsilWorks, Inc."
	description := "Sample service description"
	version := "1.0"

	mapValue(&identity, "id", &id)
	mapValue(&identity, "name", &name)
	mapValue(&identity, "copyright", &copyright)
	mapValue(&identity, "description", &description)
	mapValue(&identity, "version", &version)

	return &ServiceCore{
		id:          id,
		name:        name,
		copyright:   copyright,
		description: description,
		version:     version,
		messages:    *ln.NewLiveNote(""),
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

func (sc *ServiceCore) Get() *ServiceGetter {
	return &ServiceGetter{
		&sc.setter,
	}
}

func (sc *ServiceCore) RunMode() bool {
	return sc.productionMode
}

func (sc *ServiceCore) HostPort() int {
	return sc.hostPort
}

func mapValue[T any](identity *map[string]any, key string, out *T) {
	if identity == nil || len(*identity) == 0 {
		return
	}

	val, ok := (*identity)[key]
	if !ok {
		return
	}

	if rval, ok := val.(T); ok {
		*out = rval
	}
}

func (s *ServiceCore) Serve() {

	var (
		err error
	)

	s.modInstance = ksuid.New().String()
	s.started = time.Now()

	s.messages.AddAppMsg(s.name)
	if s.description != "" {
		s.messages.AddAppMsg(s.description)
	}
	s.messages.AddAppMsg(fmt.Sprintf(`Version %s`, s.version))
	s.messages.AddAppMsg(s.copyright)

	// start logging
	lgr, err := s.Get().Logger()
	if err != nil {
		s.messages.AddAppMsg(fmt.Sprintf(`Logger error: %s, using standard output.`, err))
	}

	s.messages.AddInfo(fmt.Sprintf(`Application id: %s`, s.id))
	s.messages.AddInfo(fmt.Sprintf(`Module instance %s`, s.modInstance))

	// succeeding functions to be initialized here

	// messages after http serve
	// serve will be run using go()

	// display logs
	for _, m := range s.messages.Notes() {
		if lgr != nil {
			lgr.Log(ifcs.LogType(m.Type), m.Message)
			continue
		}

		fmt.Printf("%s %s\r\n", time.Now().Format(time.RFC3339), m.ToString())
	}

}
