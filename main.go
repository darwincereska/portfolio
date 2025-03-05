package main

import (
	"portfolio/backend"
	"flag"
)

type FlagOptions struct {
	HOST string
	PORT int
}

func ParseFlags() *FlagOptions{
	fo := &FlagOptions{}
	flag.StringVar(&fo.HOST, "host", "localhost", "Sets the host to serve on default(localhost)")
	flag.IntVar(&fo.PORT, "port", 9999, "Sets the port default(9999)")
	flag.Parse()
	
	return fo
	
}

func main() {
	flags := ParseFlags()
	backend.ParseRoutes()
	backend.Serve(flags.HOST, flags.PORT)
}
