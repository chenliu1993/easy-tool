package main

import (
	"github.com/chenliu1993/easy-tool/pkg/cmd"
	"github.com/golang/glog"
)

func main() {
	cmd.Execute()
	glog.Flush()
}
