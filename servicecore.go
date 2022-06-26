package servicecore

import (
	"net/http"
	"sync"
	"time"
)

type ServiceCore struct {
	ID                 string             // A convenient ID for the service
	Name               string             // Name of the service
	Copyright          string             // Copyright information of the service
	Description        string             // Description of the service
	DefaultContentType string             // Default content type of the service
	DefaultHostPort    int                //
	Cache              ICache             // Caching interface
	Logger             ILogger            // Logging interface
	Queue              IQueue             //
	Data IData
	LocalDB            ILocalData         // Local database provider
	Production         bool               // Production mode flag
	Router             http.Handler       // Built-in router of the service
	Settings           *cfg.Configuration // Configuration settings of the service
	ShutdownEvent      func()             // Shutdown event processing for graceful exit
	Task               *tsk.ITask         // Task interface for service-wide task management
	Version            string             // Version of the service
	appControllerURL   string             //
	appInstance        string             //
	corsHeaders        []string           // Allowed headers for Cross Origin Resource Sharing (CORS)
	corsMethods        []string           // Allowed methods for Cross Origin Resource Sharing (CORS)
	httpServer         *http.Server       //
	launchArgs         map[string]any     // Arguments upon launch
	messagecnt         int                //
	messages           std.MessageManager // Message logged while launching the service
	mime               []std.NameValue    //
	modInstance        string             // Module instance id
	publishedEvents    []string           // List of the public events that the service publishes
	srvWG              *sync.WaitGroup    //
	started            time.Time          //
}

func Create() (*ServiceCore, error) {

}

func (sc *ServiceCore) AddCorsHeader(header string) {

}

func (sc *ServiceCore) AddCorsMethod(method string) {

}

func (sc *ServiceCore) AddPublishedEvent(event string) {

}

func (sc *ServiceCore) AddMime(event string) {

}

func (sc *ServiceCore)