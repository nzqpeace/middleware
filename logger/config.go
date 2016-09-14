package logger

import "github.com/imdario/mergo"

// Config are the options of the logger middlweare
// contains 5 bools
// Status, IP, Method, Path, EnableColors
// if set to true then these will print
type Config struct {
	// Status displays status code (bool)
	Status bool
	// IP displays request's remote address (bool)
	IP bool
	// Method displays the http method (bool)
	Method bool
	// Path displays the request path (bool)
	Path bool
	// Date displays time of request (bool)
	Date bool
	// RequestID display uuid of request (bool)
	RequestID bool
	// Body display body of request (bool)
	Body bool
	// MaxLenToPrint display max length of body
	MaxLenToPrint int
}

// DefaultConfig returns an options which all properties are true except EnableColors
func DefaultConfig() Config {
	return Config{true, true, true, true, false, true, true, 2048}
}

// Merge merges the default with the given config and returns the result
func (c Config) Merge(cfg []Config) (config Config) {

	if len(cfg) > 0 {
		config = cfg[0]
		mergo.Merge(&config, c)
	} else {
		_default := c
		config = _default
	}

	return
}
