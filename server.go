package jantra

// Server is the jantra server
type Server struct {
	cfg *Config
}

// NewServer returns a new jantra Server
func NewServer(opt ...Option) (*Server, error) {
	s := &Server{
		cfg: &Config{},
	}
	for _, o := range opt {
		o.Apply(s)
	}
	return s, nil
}
