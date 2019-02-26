package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	h bool
	v bool
	V bool
	t bool
	T bool
	q *bool

	s string
	S string
	p string
	c string
	g string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")

	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and configure options then exit")

	flag.BoolVar(&t, "t", false, "test configuration and exit")
	flag.BoolVar(&T, "T", false, "test configuration, dump it and exit")

	q = flag.Bool("q", false, "suppress non-error message during configuration testing")

	flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	flag.StringVar(&S, "S", "/usr/local/nginx", "set `prefix` path")
	flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
	flag.StringVar(&g, "g", "conf/nginx.conf", "set global `directives` out of configuration file")

	flag.Usage = useage
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
	}
	switch {
	case v:
		fmt.Printf("GET v from env:%t\n", v)
	case V:
		fmt.Printf("GET V from env:%t\n", V)
	case s != "":
		fmt.Printf("GET s from env:%s\n", s)
	}
}

func useage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
	Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]
	
	Options:
	`)
	flag.PrintDefaults()
}
