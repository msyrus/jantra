package jantra

import (
	"github.com/msyrus/jantra/backend"
	"github.com/msyrus/jantra/log"
)

// Option is the option interface of Server
type Option interface {
	Apply(*Server)
}

// OptionFunc is an implemnetation of Option interface
type OptionFunc func(*Server)

// Apply is called to apply the option
func (o OptionFunc) Apply(s *Server) {
	o(s)
}

// SetBackend option sets backend of the server
func SetBackend(b *backend.AMQP) OptionFunc {
	return func(s *Server) {
		s.back = b
	}
}

// SetLogger option sets logger of the server
func SetLogger(l log.Logger) OptionFunc {
	return func(s *Server) {
		s.logr = l
		if s.errLogr == nil {
			s.errLogr = l
		}
	}
}

// SetErrorLogger option sets error logger of the server
func SetErrorLogger(l log.Logger) OptionFunc {
	return func(s *Server) {
		s.errLogr = l
	}
}
