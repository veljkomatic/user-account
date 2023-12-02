package log

// Config is used to initialize a logger.
type Config struct {
	// Level is used to declare a level of logging.
	// Valid values are log.Debug, log.Info, log.Warn and log.Error
	Level string `json:"level"  yaml:"level"`

	// Name of the executable using the logger.
	ServiceName string

	// Service release version
	Version string
}
