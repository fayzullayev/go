package main

import "fmt"

type OptsFunc func(opts *Opts)

type Opts struct {
	maxConnections int
	id             string
	tls            bool
}

type Server struct {
	Opts Opts
}

func defaultOpts() Opts {
	return Opts{
		maxConnections: 10,
		id:             "localhost",
		tls:            false,
	}
}

func newServer(opts ...OptsFunc) *Server {
	o := defaultOpts()

	for _, opt := range opts {
		opt(&o)
	}

	return &Server{
		Opts: o,
	}
}

func setId(id string) OptsFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

func main() {
	s := newServer(func(opts *Opts) {
		opts.maxConnections = 1000
	}, setId("qwertyuio"))

	fmt.Printf("%+v\n", s)
	//myApp, err := app.NewApp()
	//if err != nil {
	//	log.Fatal("Failed to create app: ", err)
	//}
	//
	//err = myApp.Run()
	//if err != nil {
	//	log.Fatal("Error starting server: ", err)
	//}
}
