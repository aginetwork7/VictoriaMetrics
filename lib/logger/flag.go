package logger

import (
	"flag"
	"strings"

	"github.com/aginetwork7/VictoriaMetrics/lib/buildinfo"
	"github.com/aginetwork7/VictoriaMetrics/lib/flagutil"
)

func logAllFlags() {
	Infof("build version: %s", buildinfo.Version)
	Infof("command-line flags")
	flag.Visit(func(f *flag.Flag) {
		lname := strings.ToLower(f.Name)
		value := f.Value.String()
		if flagutil.IsSecretFlag(lname) {
			value = "secret"
		}
		Infof("  -%s=%q", f.Name, value)
	})
}
