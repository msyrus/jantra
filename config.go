package jantra

// Config is the jantra config
type Config struct {
}

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
