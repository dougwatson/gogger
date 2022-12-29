package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	glog.Error("hello world")
}
