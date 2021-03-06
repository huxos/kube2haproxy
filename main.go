package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/huxos/kube2dyups/app"
	"github.com/huxos/kube2dyups/app/options"

	"k8s.io/apiserver/pkg/util/logs"
	"k8s.io/apiserver/pkg/util/flag"

	"github.com/spf13/pflag"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	config := options.NewProxyServerConfig()
	config.AddFlags(pflag.CommandLine)

	flag.InitFlags()
	logs.InitLogs()
	defer logs.FlushLogs()

	s, err := app.NewProxyServerDefault(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err = s.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
