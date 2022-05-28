package servicecore

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type ServiceCore struct {
	ID          string // A convenient ID for the service
	Name        string // Name of the service
	Copyright   string // Copyright information of the service
	Description string // Description of the service

	//Args             map[string]ArgumentValue // Arguments upon launch
	AllowedHeaders []string // Allowed headers for Cross Origin Resource Sharing (CORS)
	AllowedMethods []string // Allowed methods for Cross Origin Resource Sharing (CORS)
	//BufferPool       *bpool.BufferPool        // Buffer pool for byte buffers
	Cache       ICache  // Caching interface
	Logger      ILogger // Logging interface
	ContentType string  // Default content type of the service

	LastMessageCount int // Indicates the last message count used for enumerating latest messages
	//LocalDB          *ldb.LokalDB             // Local database provider
	ProductionMode  bool     // Production mode flag
	PublishedEvents []string // List of the public events that the service publishes
	//Queue            *nats.Conn               // Built-in streaming queue client of the service
	//Router *mux.Router // Built-in router of the service
	//Settings         *cfg.Configuration       // Configuration settings of the service
	ShutdownEvent func() // Shutdown event processing for graceful exit
	//Task             tsk.ITask                // Task interface for service-wide task management
	Version          string // Version of the service
	appControllerURL string
	appInstance      string
	modInstance      string // Module instance id
	hostPort         int
	httpServer       *http.Server
	//messages         std.MessageManager // Message logged while launching the service
	messagecnt int
	//mime             []std.NameValue
	srvWG      *sync.WaitGroup
	started    time.Time
	publishing bool
	lgr        *log.Logger
}
