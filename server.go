package jantra

import (
	"github.com/msyrus/jantra/backend"
	"github.com/msyrus/jantra/log"
)

// Server is the jantra server
type Server struct {
	back          *backend.AMQP
	logr, errLogr log.Logger
	cfg           *Config
}

// NewServer returns a new jantra Server
func NewServer(opt ...Option) (*Server, error) {
	s := &Server{
		cfg: &Config{},
	}
	for _, o := range opt {
		o.Apply(s)
	}
	if s.back == nil {
		return nil, NewError("backend missing")
	}
	return s, nil
}

func print(l log.Logger, s string) {
	if l != nil {
		l.Print(s)
	}
}

func println(l log.Logger, s string) {
	if l != nil {
		l.Println(s)
	}
}

func printf(l log.Logger, fmt string, vars ...interface{}) {
	if l != nil {
		l.Printf(fmt, vars...)
	}
}
