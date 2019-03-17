package main

import (
	"github.com/mariadesouza/applog"
)

func main() {
	applog.Trace("start")
	a := []string{"words", "matter"}
	for _, s := range a {
		applog.Infof("%s\n", s)
	}
	applog.Warnf("not sure whats supposed to happen \n")
	applog.Error("this is bad")
	applog.Info("done")
}
