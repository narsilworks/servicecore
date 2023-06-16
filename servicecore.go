package servicecore

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/jessevdk/go-flags"
	ln "github.com/narsilworks/livenote"
	"github.com/narsilworks/servicecore/bare"
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

	logFile    string
	OnShutdown func() // Shutdown event processing for graceful exit

	appControllerURL string //

	httpServer *http.Server //
	opts       options      // Arguments upon launch
	messages   ln.LiveNote  // Message logged while launching the service

	appInstance     string    //
	modInstance     string    // Module instance id
	publishedEvents []string  // List of the public events that the service publishes
	started         time.Time //
	productionMode  bool      // Indicates if the service is in production mode or not
	hostPort        int       // the port where the service listens

}

func Create(identity map[string]any) (*ServiceCore, error) {

	var (
		err        error
		defLogFile string
		lOpts      options
	)

	id := "SERVICECORE"
	name := "ServiceCore Service"
	copyright := "Copyright 2023, NarsilWorks, Inc."
	description := ""
	version := "1.0"

	mapValue(&identity, "id", &id)
	mapValue(&identity, "name", &name)
	mapValue(&identity, "copyright", &copyright)
	mapValue(&identity, "description", &description)
	mapValue(&identity, "version", &version)

	// Command line parameters
	// Functions such as logging needs to know
	// the file name before serving
	flags.NewParser(&lOpts, flags.IgnoreUnknown).Parse()

	// Log File:
	// - The default log file comes from the current directory where
	//	 our application is running and affixes its id.
	// - Overrides default log file and returns it for the service starts
	if defLogFile, err = os.Getwd(); err != nil {
		defLogFile = os.TempDir()
	}
	defLogFile = path.Join(strings.ReplaceAll(defLogFile, `\`, `/`), strings.ToLower(id)+".log")
	if !isNullOrEmpty(lOpts.LogFile) {
		defLogFile = *lOpts.LogFile
	}

	return &ServiceCore{
		id:          id,
		name:        name,
		copyright:   copyright,
		description: description,
		version:     version,
		messages:    *ln.NewLiveNote(""),
		logFile:     defLogFile,
		OnShutdown:  func() {},
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

func (sc *ServiceCore) LogFile() string {
	return sc.logFile
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

func (sc *ServiceCore) Production() bool {
	return sc.productionMode
}

func (sc *ServiceCore) HostPort() int {
	return sc.hostPort
}

// Serve runs the server
func (s *ServiceCore) Serve() {

	var (
		err error
		dsl bool
	)

	// Initial values
	s.modInstance = ksuid.New().String()
	s.started = time.Now()

	s.messages.AddAppMsg(s.name)
	if s.description != "" {
		s.messages.AddAppMsg(s.description)
	}
	s.messages.AddAppMsg(fmt.Sprintf(`Version %s`, s.version))
	s.messages.AddAppMsg(s.copyright)

	// Get command line parameters
	flags.NewParser(&s.opts, flags.IgnoreUnknown).Parse()
	dsl = val(s.opts.DisableSetterLog)

	// Logger:
	// - This can be set in the main program
	// - If not set, will use the standard output
	// - Failure to load will only get a warning
	lgr, _ := s.Get().Logger()
	if lgr == nil {
		lgr = &bare.StdOutLog{}
		s.Set().Logger(lgr)
		s.messages.AddWarning(`Logger is not set. Will use standard output instead.`)
	}

	// Configuration:
	// - This can be set in the main program
	// - This can get its configuration via a file or a url as argument
	// - Failure to load will only get a warning
	stg, _ := s.Get().Config()
	if !isNullOrEmpty(s.opts.ConfigFile) {
		stg = *(new(ifcs.IConfiguration))
		if err = loadConfig(*s.opts.ConfigFile, &stg); err != nil {
			if !dsl {
				s.messages.AddWarning(fmt.Sprintf(`Configuration override failed: %s`, err))
			}
		}
	}
	if stg == nil {
		if !dsl {
			s.messages.AddWarning(`Configuration is not set.`)
		}
	} else {
		s.Set().Config(stg)
	}

	// Local Data:
	// - This can be set in the main program
	// - Failure to load will only get an information message
	ldt, _ := s.Get().LocalData()
	if ldt == nil && !dsl {
		s.messages.AddInfo(`Local Data is not set.`)
	}

	// Cache:
	// - This can be set in the main program
	// - Failure to load will only get an information message
	csh, _ := s.Get().Cache()
	if csh == nil && !dsl {
		s.messages.AddInfo(`Cache is not set. You can use github.com/eaglebush/cacheman as a memory cache.`)
	}

	// Router:
	// - This can be set in the main program
	// - Failure to load will only get an information message
	rtr, _ := s.Get().Router()
	if rtr == nil && !dsl {
		s.messages.AddInfo(`Router is not set. Your application will exit after calling Serve().`)
	}

	// CORS:
	// - This can be set in the main program
	// - Failure to load will only get an information message
	cor, _ := s.Get().CORS()
	if cor == nil && !dsl {
		if rtr != nil {
			s.messages.AddWarning(`CORS is not set. It is recommended to set this to use for routing.`)
		} else {
			s.messages.AddInfo(`CORS is not set.`)
		}
	}

	// Queue:
	// - This can be set in the main program
	// - Failure to load will only get an information message
	que, _ := s.Get().Queue()
	if que == nil && !dsl {
		s.messages.AddInfo(`Queue is not set.`)
	}

	// Data:
	// - This can be set in the main program
	// - Failure to load will only get an information message
	data, _ := s.Get().Data()
	if data == nil && !dsl {
		s.messages.AddInfo(`Data is not set.`)
	}

	// Notification:
	// - This can be set in the main program
	// - Failure to load will only get an information message
	nfs, _ := s.Get().Notification()
	if nfs == nil && !dsl {
		s.messages.AddWarning(`Notification is not set.`)
	}

	s.messages.AddInfo(fmt.Sprintf(`Application id: %s`, s.id))
	s.messages.AddInfo(fmt.Sprintf(`Module instance: %s`, s.modInstance))
	s.messages.AddInfo(fmt.Sprintf(`Service started at: %s`, s.started.Format(time.RFC3339)))

	// succeeding functions to be initialized here

	// messages after http serve
	// serve will be run using go()

	// display logs
	for _, m := range s.messages.Notes() {
		lgr.Log(ifcs.LogType(m.Type), m.Message)
	}

}
