// package AntGG is for testing only
package tools

type Opts struct {
	maxConn int
}

type Server struct {
	Opts
}

type optFunc func(*Opts)

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "defaults",
		til:     false,
	}
}

func NewServer(Opts ...optFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}
