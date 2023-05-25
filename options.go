package servicecore

// Command line options
type options struct {

	// ConfigFile - Specifies program configuration.
	ConfigFile *string `long:"config-file" short:"c"`

	// LogFile  - Specifies log file"
	LogFile *string `long:"log-file" short:"l"`

	// LocalDBFile  - Specifies local database file"
	LocalDBFile *string `long:"local-db" short:"d"`

	// HostPort - Number. Assign port. If not set, the service will use the port defined in the configuration
	HostPort *int `long:"port" short:"p"`

	// OverrideFlag - String, options are set in a url-encoded format.
	// Flags: 	n/N - notification function on or off
	//			q/Q - queue function on or off
	//			c/C - cache function on or off
	// Note:	Any flag indicated will override the start-up configuration
	// values. If ommitted, no override is activated for the specified flag.
	//
	// Example:	--override-flag=n=1&Q=off&c=on
	OverrideFlag *bool `long:"override-flag" short:"o"`

	// AppInstance - String. Application Instance ID that who launched this server
	AppInstance *string `long:"app-instance" short:"i"`

	// AppController - String. URL of the one that invoked the tasker to launch
	AppController *string `long:"app-controller" short:"a"`
}
